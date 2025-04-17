package types

import "github.com/google/uuid"

type Credentials struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt string    `json:"created_at"`
	LastLogin string    `json:"last_login"`
}

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	RocketPoints int       `json:"rocket_points"`
}

type Settigns struct {
	ID           uuid.UUID `json:"id"`
	UserId       uuid.UUID `json:"user_id"`
	ProfileImage string    `json:"profile_image"`
	StepGoal     int       `json:"step_goal"`
}
