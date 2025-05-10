package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"rocket-backend/internal/database"
)

type Server struct {
	port      int
	db        database.Service
	jwtSecret string
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	jwtSecret := os.Getenv("JWT_SECRET")

	NewServer := &Server{
		port:      port,
		db:        database.New(),
		jwtSecret: jwtSecret,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func NewServerWithDB(db database.Service, port int, jwtSecret string) *Server {
	return &Server{
		port:      port,
		db:        db, // Inject the passed DB implementation.
		jwtSecret: jwtSecret,
	}
}
