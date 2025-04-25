package server

import (
	"net/http"
	"strings"

	"rocket-backend/pkg/auth"

	"github.com/gin-gonic/gin"
)

func (s *Server) AuthMiddleware() gin.HandlerFunc {
	authService := auth.NewAuthService(s.jwtSecret)

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
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
