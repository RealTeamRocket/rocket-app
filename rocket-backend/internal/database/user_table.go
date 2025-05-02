package database

import (
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

func (s *service) SaveUserProfile(user types.User) error {
	// Assuming user.ID is the UUID linked to the credentials table
	query := `INSERT INTO users (id, username, email, rocketpoints) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, user.ID, user.Username, user.Email, user.RocketPoints)
	if err != nil {
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
	}
	return nil
}

func (s *service) GetUserByID(userID uuid.UUID) (types.User, error) {
	var user types.User
	query := `SELECT id, username, email, rocketpoints FROM users WHERE id = $1`
	err := s.db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.RocketPoints)
	if err != nil {
		return user, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}
	return user, nil
}

func (s *service) UpdateRocketPoints(userID uuid.UUID, rocketPoints int) error {
	query := `UPDATE users SET rocketpoints = rocketpoints + $2 WHERE id = $1`
	_, err := s.db.Exec(query, userID, rocketPoints)
	if err != nil {
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToUpdate, err)
	}
	return nil
}

func (s *service) GetUserIDByName(name string) (uuid.UUID, error) {
 var userID uuid.UUID
 query := `SELECT id FROM users WHERE username = $1`
 err := s.db.QueryRow(query, name).Scan(&userID)
 if err != nil {
  return uuid.Nil, fmt.Errorf("failed to get user ID by name: %w", err)
 }
 return userID, nil
}
