package server

import (
	"net/http"
	"strings"
	"os"

	"rocket-backend/pkg/auth"

	_ "github.com/joho/godotenv/autoload"
	"github.com/gin-gonic/gin"
)

func (s *Server) AuthMiddleware() gin.HandlerFunc {
	authService := auth.NewAuthService(s.jwtSecret)

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var tokenString string

		if authHeader != "" {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			cookie, err := c.Cookie("jwt_token")
			if err != nil || cookie == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header or jwt_token cookie required"})
				c.Abort()
				return
			}
			tokenString = cookie
		}

		token, err := authService.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userUUID, err := authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Check if the user exists in the database
		_, err = s.db.GetUserByID(userUUID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist"})
			c.Abort()
			return
		}

		c.Set("userID", userUUID.String())
		c.Next()
	}
}

func (s *Server) APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedAPIKey := os.Getenv("API_KEY")
		sentApiKey := c.GetHeader("X-API-KEY")

		// Debug-Logs
		if expectedAPIKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not configured"})
			c.Abort()
			return
		}

		if sentApiKey != expectedAPIKey || sentApiKey == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Invalid or missing API key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
