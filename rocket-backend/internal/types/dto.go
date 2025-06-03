package types

import (
	"time"

	"github.com/google/uuid"
)

type RegisterDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateStepsDTO struct {
	Steps int `json:"steps"`
}

type SettingsDTO struct {
	StepGoal int `json:"stepGoal"`
}

type GetImageDTO struct {
	UserID string `json:"user_id" binding:"required,uuid"`
}

type CompleteChallengesDTO struct {
	ChallengeID  uuid.UUID `json:"challenge_id" binding:"required"`
	RocketPoints int       `json:"rocket_points" binding:"required"`
}

type StepStatistic struct {
	Day   string `json:"day"`
	Steps int    `json:"steps"`
}

type UserWithImageDTO struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	RocketPoints int       `json:"rocket_points"`
	ImageName    string    `json:"image_name"`
	ImageData    string    `json:"image_data"`
}

type RunDataDTO struct {
    Route    string  `json:"route" binding:"required"`    // WKT LineString
    Duration string  `json:"duration" binding:"required"`
    Distance float64 `json:"distance"`
}

type RunDTO struct {
    ID       string  `json:"id"`
    Route    string  `json:"route"`
    Duration string  `json:"duration"`
    Distance float64 `json:"distance"`
    CreatedAt string `json:"created_at"`
}

type ActivityWithUser struct {
	Name    string    `json:"name"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

type ChatMessage struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	Timestamp string    `json:"timestamp"`
	Reactions int       `json:"reactions"`
	HasReacted bool     `json:"hasReacted"`
}
