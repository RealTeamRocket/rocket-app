package database

import (
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

func (s *service) SaveCredentials(creds types.Credentials) error {
	query := `INSERT INTO credentials (id, email, password, created_at, last_login) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, creds.ID, creds.Email, creds.Password, creds.CreatedAt, creds.LastLogin)
	if err != nil {
		logger.Error("Failed to save credentials", err)
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
	}
	return nil
}

func (s *service) GetUserByEmail(email string) (types.Credentials, error) {
	var creds types.Credentials
	query := `SELECT id, email, password, created_at, last_login FROM credentials WHERE email = $1`
	err := s.db.QueryRow(query, email).Scan(&creds.ID, &creds.Email, &creds.Password, &creds.CreatedAt, &creds.LastLogin)
	if err != nil {
		logger.Error("Failed to get user by email", err)
		return creds, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}
	return creds, nil
}

func (s *service) CheckEmail(email string) error {
	query := `SELECT COUNT(*) FROM credentials WHERE email = $1`
	var count int
	err := s.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		logger.Error("Failed to check email", err)
		return fmt.Errorf("%w: %v", custom_error.ErrDatabaseQuery, err)
	}
	if count > 0 {
		return custom_error.ErrEmailAlreadyExists
	}
	return nil
}
