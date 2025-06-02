package types

import (
	"github.com/google/uuid"
	"time"
)

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

type Settings struct {
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	ImageId  uuid.UUID `json:"image_id"`
	StepGoal int       `json:"step_goal"`
}

type UserImage struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Data []byte    `json:"data"`
}

type Activity struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}
