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

func (s *service) GetAllUsers() ([]types.User, error) {
	var users []types.User
	query := `SELECT id, username, email, rocketpoints FROM users`
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	defer rows.Close()

	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.RocketPoints); err != nil {
			return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	return users, nil
}
