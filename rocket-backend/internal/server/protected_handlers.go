package server

import (
	"net/http"
	"rocket-backend/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) AuthHelloHandler(c *gin.Context) {
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

	cred, err := s.db.GetUserByID(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hello from auth", "user": cred.Username})
}

func (s *Server) Authenticated(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"authenticated": "true"})
}

func (s *Server) UpdateSteps(c *gin.Context) {
	var updateStepDTO types.UpdateStepsDTO
	if err := c.ShouldBindJSON(&updateStepDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

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

	if err := s.db.UpdateDailySteps(userUUID, updateStepDTO.Steps); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong in db"})
		return
	}

	c.JSON(http.StatusOK,  gin.H{"message": "Daily Steps saved"})
}
