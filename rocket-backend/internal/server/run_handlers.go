package server

import (
	"fmt"
	"net/http"

	"rocket-backend/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

	message := "Completed a " + fmt.Sprintf("%.2f", runData.Distance) + " km run in " + runData.Duration + " minutes"
	err = s.db.SaveActivity(userUUID, message)

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

func (s *Server) DeleteRunHandler(c *gin.Context) {
	runIDStr := c.Param("id")
	runID, err := uuid.Parse(runIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid run ID format"})
		return
	}

	err = s.db.DeleteRun(runID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete run"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Run deleted successfully"})
}

func (s *Server) PlanRunHandler(c *gin.Context) {
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

    var req struct {
        Route string `json:"route" binding:"required"`
        Name  string `json:"name" binding:"required"`
        Distance float64 `json:"distance" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    err = s.db.SavePlannedRun(userUUID, req.Route, req.Name, req.Distance)
    if err != nil {
        if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "A planned run with this name already exists."})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save planned run"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Planned run saved successfully"})
}

func (s *Server) GetPlannedRunHandler(c *gin.Context) {
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

    runs, err := s.db.GetAllPlannedRunsByUser(userUUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch planned runs"})
        return
    }

    c.JSON(http.StatusOK, runs)
}
