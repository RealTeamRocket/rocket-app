package server_tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"rocket-backend/internal/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/gorilla/websocket"
)

var _ = Describe("Chat WebSocket API", func() {
	var token string

	BeforeEach(func() {
		token = registerAndLogin("wsuser@example.com", "password123", "wsuser")
	})

	It("should send and receive chat messages over websocket", func() {
		// Build ws URL (convert http://localhost:8090/api/v1/protected/ws/chat to ws://...)
		httpURL := baseURL + "/protected/ws/chat"
		u, err := url.Parse(httpURL)
		Expect(err).To(BeNil())
		u.Scheme = "ws"

		// Add JWT as Authorization header
		header := http.Header{}
		header.Set("Authorization", "Bearer "+token)

		// Connect to websocket
		ws, resp, err := websocket.DefaultDialer.Dial(u.String(), header)
		Expect(err).To(BeNil(), fmt.Sprintf("WebSocket dial failed: %v (HTTP status: %d)", err, resp.StatusCode))
		defer ws.Close()

		// Send a chat message
		msg := map[string]string{"message": "hello from integration test"}
		msgBytes, _ := json.Marshal(msg)
		err = ws.WriteMessage(websocket.TextMessage, msgBytes)
		Expect(err).To(BeNil())

		// Receive the broadcasted message
		ws.SetReadDeadline(time.Now().Add(2 * time.Second))
		_, respBytes, err := ws.ReadMessage()
		Expect(err).To(BeNil())

		var chatMsg types.ChatMessage
		err = json.Unmarshal(respBytes, &chatMsg)
		Expect(err).To(BeNil())
		Expect(chatMsg.Message).To(Equal("hello from integration test"))
		Expect(chatMsg.Username).To(Equal("wsuser"))
	})
})
