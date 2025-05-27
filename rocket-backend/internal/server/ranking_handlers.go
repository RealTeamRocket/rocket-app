package server

import (
	"errors"
	"net/http"

	"rocket-backend/internal/custom_error"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

func (s *Server) GetUserRanking(c *gin.Context) {
	ranking, err := s.db.GetTopUsers(100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching ranking"})
		return
	}
	if ranking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ranking for user not found"})
		return
	}

	c.JSON(http.StatusOK, ranking)
}
