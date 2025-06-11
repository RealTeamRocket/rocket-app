package server

import (
	"encoding/base64"
	"errors"
	"net/http"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) UpdateStepsHandler(c *gin.Context) {
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

func (s *Server) GetUserStatisticsHandler(c *gin.Context) {
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

func (s *Server) GetUserImageHandler(c *gin.Context) {
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
		// Proceed without returning an error, as the image might not exist
		img = nil
	}

	user, err := s.db.GetUserByID(userUUID)
	if err != nil {
		logger.Error("Failed to get user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	response := gin.H{
		"username":  user.Username,
		"name":      nil,
		"mime_type": nil,
		"data":      nil,
	}

	if img != nil {
		encodedImage := base64.StdEncoding.EncodeToString(img.Data)
		response["name"] = img.Name
		response["mime_type"] = http.DetectContentType(img.Data)
		response["data"] = encodedImage
	}

	c.JSON(http.StatusOK, response)
}

func (s *Server) GetRocketPointsHandler(c *gin.Context) {
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

	rocketPoints, err := s.db.GetRocketPointsByUserID(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve rocket points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rocket_points": rocketPoints})
}

func (s *Server) GetAllUsersHandler(c *gin.Context) {
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

	users, err := s.db.GetAllUsers(&userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	var usersWithImages []types.UserWithImageDTO

	for _, user := range users {
		var userWithImage types.UserWithImageDTO
		userWithImage.ID = user.ID
		userWithImage.Username = user.Username
		userWithImage.Email = user.Email
		userWithImage.RocketPoints = user.RocketPoints
		userImage, err := s.db.GetUserImage(user.ID)

		if err != nil {
			logger.Warn("Failed to fetch image for user %s: %v\n", user.ID, err)
			userWithImage.ImageName = ""
			userWithImage.ImageData = ""
		} else if userImage != nil {
			userWithImage.ImageName = userImage.Name
			userWithImage.ImageData = base64.StdEncoding.EncodeToString(userImage.Data)
		}

		usersWithImages = append(usersWithImages, userWithImage)
	}
	c.JSON(http.StatusOK, usersWithImages)
}

func (s *Server) GetUserHandler(c *gin.Context) {
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

	user, err := s.db.GetUserByID(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) GetUserByNameHandler(c *gin.Context) {
	username := c.Param("name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	userID, err := s.db.GetUserIDByName(username)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		}
		return
	}

	user, err := s.db.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		}
		return
	}

	userImage, err := s.db.GetUserImage(userID)
	var imageName, imageData string
	if err != nil || userImage == nil {
		imageName = ""
		imageData = ""
	} else {
		imageName = userImage.Name
		imageData = base64.StdEncoding.EncodeToString(userImage.Data)
	}

	userWithImage := types.UserWithImageDTO{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		RocketPoints: user.RocketPoints,
		ImageName:    imageName,
		ImageData:    imageData,
	}

	c.JSON(http.StatusOK, userWithImage)
}
