package server

import (
	"net/http"

	"rocket-backend/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) UploadRunHandler(c *gin.Context) {
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

	var runData types.RunDataDTO
	if err := c.ShouldBindJSON(&runData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = s.db.SaveRun(userUUID, runData.Route, runData.Duration, runData.Distance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save run"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Run data uploaded successfully"})
}

func (s *Server) GetAllRunsHandler(c *gin.Context) {
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

	runs, err := s.db.GetAllRunsByUser(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch runs"})
		return
	}

	c.JSON(http.StatusOK, runs)
}
