package server

import (
	"net/http"
	"rocket-backend/internal/auth"
	"rocket-backend/internal/types"
	"time"

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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	authService := auth.NewAuthService(s.jwtSecret)
	tokenString, err := authService.GenerateToken(storedCreds.ID)
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
	var registerDto types.RegisterDTO
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := s.db.CheckEmail(registerDto.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	var creds types.Credentials
	creds.ID = uuid.New()
	creds.Email = registerDto.Email
	creds.Password = string(hashedPassword)
	creds.CreatedAt = time.Now().Format(time.RFC3339)
	creds.LastLogin = creds.CreatedAt

	if err := s.db.SaveCredentials(creds); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save credentials"})
		return
	}

	var user types.User
	user.ID = creds.ID
	user.Username = registerDto.Username
	user.Email = registerDto.Email
	user.RocketPoints = 0

	if err := s.db.SaveUserProfile(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
