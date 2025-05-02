package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"rocket-backend/internal/challenges"
	"rocket-backend/internal/custom_error"
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
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		}
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

	c.JSON(http.StatusOK, gin.H{"message": "Daily Steps saved"})
}

func (s *Server) UpdateSettings(c *gin.Context) {
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

	// maximus is 10mb
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	settingsJSON := c.PostForm("settings")
	var settingsDTO types.SettingsDTO
	if err := json.Unmarshal([]byte(settingsJSON), &settingsDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid settings format"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image"})
		return
	}

	var imageID uuid.UUID
	if err == nil {
		defer file.Close()

		// Read image into bytes
		imageData, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image data"})
			return
		}

		// Check the content type
		mimeType := http.DetectContentType(imageData)
		if mimeType != "image/jpeg" && mimeType != "image/png" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Only JPEG (.jpg, .jpeg) and PNG (.png) images are allowed"})
			return
		}

		// Save image
		imageID, err = s.db.SaveImage(header.Filename, imageData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
	}

	// Save settings
	err = s.db.UpdateSettings(userUUID, settingsDTO, imageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update settings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully"})
}

func (s *Server) GetSettings(c *gin.Context) {
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

func (s *Server) AddFriend(c *gin.Context) {
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

	friendName := c.PostForm("friend_name")
	if friendName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "friend_name is required"})
		return
	}

	friendID, err := s.db.GetUserIDByName(friendName)
	if err != nil {
		// TODO: better error handling here need to wait for #46
		c.JSON(http.StatusBadRequest, gin.H{"error": "friend_name is required"})
		return
	}

	err = s.db.AddFriend(userUUID, friendID)
	if err != nil {
		// TODO: better error handling here need to wait for #46
		c.JSON(http.StatusBadRequest, gin.H{"error": "friend_name is required"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "friend added successful"})
}

func (s *Server) GetAllFriends(c *gin.Context) {
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

	friends, err := s.db.GetFriends(userUUID)

	c.JSON(http.StatusLocked, friends)
}

func (s *Server) GetFriendsRanked(c *gin.Context) {
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

	friends, err := s.db.GetFriendsRankedByPoints(userUUID)

	c.JSON(http.StatusLocked, friends)
}

func (s *Server) GetDailyChallenges(c *gin.Context) {
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

	challengeManager := challenges.NewChallengeManager(s.db)
	dailies, err := challengeManager.GetDailies(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrChallengeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not enough challenges available"})
		} else if errors.Is(err, custom_error.ErrFailedToRetrieveData) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve challenges"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, dailies)
}

func (s *Server) CompleteChallenge(c *gin.Context) {
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

	var pointsDTO types.CompleteChallengesDTO
	if err := c.ShouldBindJSON(&pointsDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = s.db.UpdateRocketPoints(userUUID, pointsDTO.RocketPoints)
	if err != nil {
		if errors.Is(err, custom_error.ErrFailedToUpdate) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rocket points"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	err = s.db.CompleteChallenge(userUUID, pointsDTO)
	if err != nil {
		if errors.Is(err, custom_error.ErrFailedToUpdate) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete challenge"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Challenge completed successfully"})
}
