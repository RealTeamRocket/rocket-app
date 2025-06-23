package server

import (
	"encoding/base64"
	"net/http"

	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) GetActivityHandler(c *gin.Context) {
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

	activities, err := s.db.GetActivitiesForUserAndFriends(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}

	user, err := s.db.GetUserByID(userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
		return
	}

	var activitiesWithImages []types.ActivityWithUserAndImage
	for _, activity := range activities {
		img, err := s.db.GetUserImage(activity.UserID)
		if err != nil {
			logger.Error("Failed to get image", err)
			img = nil
		}

		if img != nil {
			encodedImage := base64.StdEncoding.EncodeToString(img.Data)
			activitiesWithImages = append(activitiesWithImages, types.ActivityWithUserAndImage{
				Name:      activity.Name,
				Time:      activity.Time,
				Message:   activity.Message,
				ImageName: img.Name,
				ImageType: http.DetectContentType(img.Data),
				ImageData: encodedImage,
			})
		} else {
			activitiesWithImages = append(activitiesWithImages, types.ActivityWithUserAndImage{
				Name:      activity.Name,
				Time:      activity.Time,
				Message:   activity.Message,
				ImageName: "",
				ImageType: "",
				ImageData: "",
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username, "activities": activitiesWithImages})
}
