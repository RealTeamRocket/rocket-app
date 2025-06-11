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

func (s *Server) AddFriendHandler(c *gin.Context) {
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

func (s *Server) DeleteFriendHandler(c *gin.Context) {
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

	friendName := c.Param("name")
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

func (s *Server) GetAllFriendsHandler(c *gin.Context) {
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

	var friendsWithImages []types.UserWithImageDTO
	for _, fr := range friends {
		var f types.UserWithImageDTO
		f.ID = fr.ID
		f.Username = fr.Username
		f.Email = fr.Email
		f.RocketPoints = fr.RocketPoints

		userImage, imgErr := s.db.GetUserImage(fr.ID)
		if imgErr != nil {
			// logger.Warn("Failed to fetch image for friend %s: %v\n", fr.ID, imgErr)
			f.ImageName = ""
			f.ImageData = ""
		} else if userImage != nil {
			f.ImageName = userImage.Name
			f.ImageData = base64.StdEncoding.EncodeToString(userImage.Data)
		}
		friendsWithImages = append(friendsWithImages, f)
	}

	c.JSON(http.StatusOK, friendsWithImages)
}

func (s *Server) GetFollowingHandler(c *gin.Context) {
	paramID := c.Param("id")
	if paramID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	userUUID, err := uuid.Parse(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Get the list of users this user is following
	following, err := s.db.GetFriends(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": custom_error.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom_error.ErrFailedToRetrieveData.Error()})
		}
		return
	}

	var followingWithImages []types.UserWithImageDTO
	for _, fr := range following {
		var f types.UserWithImageDTO
		f.ID = fr.ID
		f.Username = fr.Username
		f.Email = fr.Email
		f.RocketPoints = fr.RocketPoints

		userImage, imgErr := s.db.GetUserImage(fr.ID)
		if imgErr != nil {
			f.ImageName = ""
			f.ImageData = ""
		} else if userImage != nil {
			f.ImageName = userImage.Name
			f.ImageData = base64.StdEncoding.EncodeToString(userImage.Data)
		}
		followingWithImages = append(followingWithImages, f)
	}

	c.JSON(http.StatusOK, followingWithImages)
}

func (s *Server) GetFollowersHandler(c *gin.Context) {
	paramID := c.Param("id")
	if paramID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	userUUID, err := uuid.Parse(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	followers, err := s.db.GetFollowers(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": custom_error.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom_error.ErrFailedToRetrieveData.Error()})
		}
		return
	}

	var followersWithImages []types.UserWithImageDTO
	for _, fr := range followers {
		var f types.UserWithImageDTO
		f.ID = fr.ID
		f.Username = fr.Username
		f.Email = fr.Email
		f.RocketPoints = fr.RocketPoints

		userImage, imgErr := s.db.GetUserImage(fr.ID)
		if imgErr != nil {
			f.ImageName = ""
			f.ImageData = ""
		} else if userImage != nil {
			f.ImageName = userImage.Name
			f.ImageData = base64.StdEncoding.EncodeToString(userImage.Data)
		}
		followersWithImages = append(followersWithImages, f)
	}

	c.JSON(http.StatusOK, followersWithImages)
}
