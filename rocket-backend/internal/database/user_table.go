package database

import (
	"database/sql"
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
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

func (s *service) GetRocketPointsByUserID(userID uuid.UUID) (int, error) {
	var rocketPoints int
	query := `SELECT rocketpoints FROM users WHERE id = $1`
	err := s.db.QueryRow(query, userID).Scan(&rocketPoints)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}
	return rocketPoints, nil
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

func (s *service) GetTopUsers(limit int) ([]types.User, error) {
	var users []types.User
	query := `SELECT id, username, email, rocketpoints FROM users ORDER BY rocketpoints DESC LIMIT $1`
	rows, err := s.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.RocketPoints); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}

// UpdateUserName updates the username for a given user UUID.
func (s *service) UpdateUserName(userID uuid.UUID, newName string) error {
	query := `UPDATE users SET username = $2 WHERE id = $1`
	_, err := s.db.Exec(query, userID, newName)
	if err != nil {
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToUpdate, err)
	}
	return nil
}

// UpdateUserEmail updates the email for a given user UUID in both users and credentials tables.
func (s *service) UpdateUserEmail(userID uuid.UUID, newEmail string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Update in users table
	queryUsers := `UPDATE users SET email = $2 WHERE id = $1`
	if _, err := tx.Exec(queryUsers, userID, newEmail); err != nil {
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToUpdate, err)
	}

	// Update in credentials table
	queryCreds := `UPDATE credentials SET email = $2 WHERE id = $1`
	if _, err := tx.Exec(queryCreds, userID, newEmail); err != nil {
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToUpdate, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

// CheckUserPassword checks if the provided password matches the user's current password using bcrypt.
func (s *service) CheckUserPassword(userID uuid.UUID, currentPassword string) (bool, error) {
	var hashedPassword string
	query := `SELECT password FROM credentials WHERE id = $1`
	err := s.db.QueryRow(query, userID).Scan(&hashedPassword)
	if err != nil {
		return false, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}
	// Use bcrypt for password verification
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currentPassword)); err != nil {
		return false, nil
	}
	return true, nil
}

// UpdateUserPassword updates the user's password in the credentials table using bcrypt.
func (s *service) UpdateUserPassword(userID uuid.UUID, newPassword string) error {
	// Hash the new password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}
	query := `UPDATE credentials SET password = $2 WHERE id = $1`
	_, err = s.db.Exec(query, userID, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToUpdate, err)
	}
	return nil
}

func (s *service) GetAllUsers(excludeUserID *uuid.UUID) ([]types.User, error) {
	var users []types.User
	var rows *sql.Rows
	var err error

	if excludeUserID != nil {
		query := `SELECT id, username, email, rocketpoints FROM users WHERE id != $1`
		rows, err = s.db.Query(query, *excludeUserID)
	} else {
		query := `SELECT id, username, email, rocketpoints FROM users`
		rows, err = s.db.Query(query)
	}

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

func (s *service) DeleteUser(userID uuid.UUID) error {
	// First, delete from users (this will cascade to all dependent tables)
	query := `DELETE FROM users WHERE id = $1`
	_, err := s.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToDelete, err)
	}

	// Then, delete from credentials (since users references credentials, not the other way around)
	credQuery := `DELETE FROM credentials WHERE id = $1`
	_, credErr := s.db.Exec(credQuery, userID)
	if credErr != nil {
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToDelete, credErr)
	}

	return nil
}
