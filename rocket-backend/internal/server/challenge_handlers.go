package server

import (
	"errors"
	"net/http"

	"rocket-backend/internal/challenges"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Assumes Server struct is defined in the server package and imported appropriately
// If not, adjust the import or struct definition as needed

func (s *Server) GetDailyChallengesHandler(c *gin.Context) {
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

func (s *Server) CompleteChallengeHandler(c *gin.Context) {
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

	// save activity
	challenge, err := s.db.GetChallengeByID(pointsDTO.ChallengeID)
	if err != nil {
		challenge.Text = "Unknown challenge"
	}

	message := "Completed a daily challenge: " + challenge.Text
	_ = s.db.SaveActivity(userUUID, message)

	c.JSON(http.StatusOK, gin.H{"message": "Challenge completed successfully"})
}

func (s *Server) InviteFriendChallenge(c *gin.Context) {
	var req struct {
		ChallengeID string `json:"challenge_id" binding:"required"`
		FriendID    string `json:"friend_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	friendUUID, err := uuid.Parse(req.FriendID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid friend ID format"})
		return
	}

	challengeUUID, err := uuid.Parse(req.ChallengeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid challenge ID format"})
		return
	}

	err = s.db.InviteFriendToChallenge(challengeUUID, friendUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrFailedToUpdate) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invite friend"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend invited successfully"})
}

func (s *Server) GetDailyChallengeProgress(c *gin.Context) {
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

	completed, err := s.db.GetCompletedChallengesAmount(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrFailedToRetrieveData) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve challenge progress"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	total, err := s.db.GetAllChallengesAmount(userUUID)
	if err != nil {
		if errors.Is(err, custom_error.ErrFailedToRetrieveData) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve total challenges"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	progress := types.DailyChallengeProgress{
		Completed: completed,
		Total:     total,
	}

	c.JSON(http.StatusOK, progress)
}
