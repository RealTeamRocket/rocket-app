package server

import (
	"io"
	"net/http"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) GetSettingsHandler(c *gin.Context) {
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

	settings, err := s.db.GetSettingsByUserID(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching settings"})
		return
	}
	if settings == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "settings for user where not found"})
		return
	}

	c.JSON(http.StatusOK, settings)
}

func (s *Server) UpdateStepGoalHandler(c *gin.Context) {
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

	var stepGoalDTO types.SettingsDTO
	if err := c.ShouldBindJSON(&stepGoalDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	logger.Debug("Received JSON body for step goal update", "stepGoalDTO", stepGoalDTO)
	logger.Debug("Received step goal update request", "userID", userUUID, "stepGoal", stepGoalDTO.StepGoal)
	logger.Debug("UpdateStepGoal handler reached")

	if stepGoalDTO.StepGoal <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Step goal must be greater than 0"})
		return
	}

	logger.Info("Updating step goal", "userID", userUUID, "stepGoal", stepGoalDTO.StepGoal)

	err = s.db.UpdateStepGoal(userUUID, stepGoalDTO.StepGoal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update step goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Step goal updated successfully"})
}

func (s *Server) UpdateImageHandler(c *gin.Context) {
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

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image"})
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image data"})
		return
	}

	if header == nil || len(imageData) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image provided"})
		return
	}

	logger.Info("Updating image", "userID", userUUID, "imageName", header.Filename)

	imageID, err := s.db.SaveImage(header.Filename, imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	err = s.db.UpdateImage(userUUID, imageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image updated successfully"})
}

func (s *Server) DeleteImageHandler(c *gin.Context) {
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

	logger.Info("Deleting image for user", "userID", userUUID)

	err = s.db.DeleteUserImage(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
