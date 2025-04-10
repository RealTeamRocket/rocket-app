package types

import "github.com/google/uuid"

type Credentials struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt string    `json:"created_at"`
	LastLogin string    `json:"last_login"`
}
