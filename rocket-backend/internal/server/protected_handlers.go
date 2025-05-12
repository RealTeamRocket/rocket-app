package server

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"rocket-backend/internal/challenges"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

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
		c.JSON(http.StatusUnauthorized, gin.H{"error": custom_error.ErrUserNotFound.Error()})
		return
	}

	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	friendName := struct {
		FriendName string `json:"friend_name"`
	}{}

	if err := c.ShouldBindJSON(&friendName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if friendName.FriendName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "friend_name is required"})
		return
	}

	friendID, err := s.db.GetUserIDByName(friendName.FriendName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": custom_error.ErrUserNotFound.Error()})
		return
	}

	err = s.db.AddFriend(userUUID, friendID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": custom_error.ErrFailedToSave.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend added successfully"})
}

func (s *Server) DeleteFriend(c *gin.Context) {
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
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Friend not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve friend information"})
		}
		return
	}

	err = s.db.DeleteFriend(userUUID, friendID)
	if err != nil {
		if errors.Is(err, custom_error.ErrFailedToDelete) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend deleted successfully"})
}

func (s *Server) GetAllFriends(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": custom_error.ErrUserNotFound.Error()})
		return
	}

	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	friends, err := s.db.GetFriends(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": custom_error.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom_error.ErrFailedToRetrieveData.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, friends)
}

func (s *Server) GetFriendsRanked(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": custom_error.ErrUserNotFound.Error()})
		return
	}

	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	friends, err := s.db.GetFriendsRankedByPoints(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": custom_error.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom_error.ErrFailedToRetrieveData.Error()})
		}
		return
	}

	if len(friends) == 0 {
		c.JSON(http.StatusOK, []interface{}{})
		return
	}

	c.JSON(http.StatusOK, friends)
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

func (s *Server) GetUserImage(c *gin.Context) {
	var req types.GetImageDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}
		req.UserID = userID.(string)
	}

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
		return
	}

	img, err := s.db.GetUserImage(userUUID)
	if err != nil {
		logger.Error("Failed to get image", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve image"})
		return
	}
	if img == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No image found for user"})
		return
	}

	// Encode image data as Base64
	encodedImage := base64.StdEncoding.EncodeToString(img.Data)

	user, err := s.db.GetUserByID(userUUID)
	if err != nil {
		logger.Error("Failed to get user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username":  user.Username,
		"name":      img.Name,
		"mime_type": http.DetectContentType(img.Data),
		"data":      encodedImage,
	})
}

func (s *Server) getUserStatistics(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user := struct {
		ID string `json:"id"`
	}{}

	if err := c.ShouldBindJSON(&user); err != nil || user.ID == "" {
		user.ID = userID.(string)
	}

	userUUID, err := uuid.Parse(user.ID)
	if err != nil {
		logger.Debug("here is the upsi")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	stats, err := s.db.GetUserStatistics(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user statistics"})
		}
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (s *Server) UpdateStepGoal(c *gin.Context) {
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

func (s *Server) UpdateImage(c *gin.Context) {
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
