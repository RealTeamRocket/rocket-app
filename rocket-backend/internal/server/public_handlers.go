package server

import (
	"net/http"
	"rocket-backend/internal/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) LoginHandler(c *gin.Context) {
	var creds types.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	storedCreds, err := s.db.GetUserByEmail(creds.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid username or password"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedCreds.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie(
		"jwt_token",
		tokenString,
		3600*72,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (s *Server) RegisterHandler(c *gin.Context) {
	var creds types.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := s.db.CheckEmail(creds.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	creds.ID = uuid.New()
	creds.Password = string(hashedPassword)
	creds.Username = creds.Username
	creds.CreatedAt = time.Now().Format(time.RFC3339)
	creds.LastLogin = creds.CreatedAt

	if err := s.db.SaveCredentials(creds); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
