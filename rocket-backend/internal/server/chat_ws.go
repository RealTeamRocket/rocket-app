package server

import (
	"encoding/json"
	"net/http"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ChatHub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.Mutex
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (hub *ChatHub) Run() {
	for {
		select {
		case conn := <-hub.register:
			hub.mu.Lock()
			hub.clients[conn] = true
			hub.mu.Unlock()
		case conn := <-hub.unregister:
			hub.mu.Lock()
			if _, ok := hub.clients[conn]; ok {
				delete(hub.clients, conn)
				conn.Close()
			}
			hub.mu.Unlock()
		case message := <-hub.broadcast:
			hub.mu.Lock()
			for conn := range hub.clients {
				err := conn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					conn.Close()
					delete(hub.clients, conn)
				}
			}
			hub.mu.Unlock()
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (s *Server) ChatWebSocketHandler(hub *ChatHub) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		userUUID, err := uuid.Parse(userID.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
			return
		}

		user, err := s.db.GetUserByID(userUUID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to find user"})
			return
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logger.Error("WebSocket upgrade error: ", err)
			return
		}
		hub.register <- conn

		defer func() {
			hub.unregister <- conn
		}()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			// Unmarshal the incoming message
			var incoming struct {
				Message string `json:"message"`
			}
			if err := json.Unmarshal(msg, &incoming); err != nil {
				logger.Error("Failed to unmarshal incoming message: " ,err)
				continue
			}
			// Create the outgoing message with username
			outgoing := types.ChatMessage{
				Username: user.Username,
				Message:  incoming.Message,
				Timestamp: time.Now().Format(time.RFC3339),
			}
			outBytes, _ := json.Marshal(outgoing)
			hub.broadcast <- outBytes
		}
	}
}
