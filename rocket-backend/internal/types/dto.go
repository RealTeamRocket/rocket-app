package types

import "github.com/google/uuid"

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
