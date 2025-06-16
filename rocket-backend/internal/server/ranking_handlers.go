package server

import (
	"encoding/base64"
	"errors"
	"net/http"

	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) GetFriendsRankedHandler(c *gin.Context) {
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

	var friendsWithImages []types.UserWithImageDTO
	for _, user := range friends {
		userImage, err := s.db.GetUserImage(user.ID)
		var imageName, imageData string
		if err == nil && userImage != nil {
			imageName = userImage.Name
			imageData = base64.StdEncoding.EncodeToString(userImage.Data)
		}
		friendsWithImages = append(friendsWithImages, types.UserWithImageDTO{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			RocketPoints: user.RocketPoints,
			ImageName:    imageName,
			ImageData:    imageData,
			Steps:        0,
		})
	}

	c.JSON(http.StatusOK, friendsWithImages)
}

func (s *Server) GetUserRankingHandler(c *gin.Context) {
	ranking, err := s.db.GetTopUsers(100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching ranking"})
		return
	}
	if ranking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ranking for user not found"})
		return
	}

	var usersWithImages []types.UserWithImageDTO
	for _, user := range ranking {
		userImage, err := s.db.GetUserImage(user.ID)
		var imageName, imageData string
		if err == nil && userImage != nil {
			imageName = userImage.Name
			imageData = base64.StdEncoding.EncodeToString(userImage.Data)
		}

		usersWithImages = append(usersWithImages, types.UserWithImageDTO{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			RocketPoints: user.RocketPoints,
			ImageName:    imageName,
			ImageData:    imageData,
			Steps:        0, // no steps needed for ranking
		})
	}

	c.JSON(http.StatusOK, usersWithImages)
}
