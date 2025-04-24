package database

import (
	"fmt"

	"rocket-backend/internal/types"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

func (s *service) SaveUserProfile(user types.User) error {
	// Assuming user.ID is the UUID linked to the credentials table
	query := `INSERT INTO users (id, username, email, rocketpoints)
	          VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, user.ID, user.Username, user.Email, user.RocketPoints)
	if err != nil {
		return fmt.Errorf("failed to save user profile: %w", err)
	}
	return nil
}

func (s *service) GetUserByID(userID uuid.UUID) (types.User, error) {
	var user types.User
	query := `SELECT id, username, email, rocketpoints FROM users WHERE id = $1`
	err := s.db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.RocketPoints)
	if err != nil {
		return user, fmt.Errorf("failed to get user by username: %w", err)
	}
	return user, nil
}

func (s *service) UpdateRocketPoints(userID uuid.UUID, rocketPoints int) error {
	query := `UPDATE users SET rocketpoints = rocketpoints + $2 WHERE id = $1`
	_, err := s.db.Exec(query, userID, rocketPoints)
	if err != nil {
		return fmt.Errorf("failed to update rocket points: %w", err)
	}
	return nil
}
