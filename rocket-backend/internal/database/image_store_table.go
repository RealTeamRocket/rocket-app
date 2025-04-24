package database

import (
	"database/sql"
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
)

func (s *service) SaveImage(filename string, data []byte) (uuid.UUID, error) {
	id := uuid.New()

	_, err := s.db.Exec(`
		INSERT INTO image_store (id, image_name, image_data)
		VALUES ($1, $2, $3)
	`, id, filename, data)

	if err != nil {
		logger.Error("Failed to save image", err)
		return uuid.Nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
	}

	return id, nil
}

func (s *service) GetUserImage(userID uuid.UUID) (*types.UserImage, error) {
	query := `
		SELECT
			i.id, i.image_name, i.image_data
		FROM
			settings s
		JOIN
			image_store i ON s.image_id = i.id
		WHERE
			s.user_id = $1
	`

	var img types.UserImage
	err := s.db.QueryRow(query, userID).Scan(
		&img.ID,
		&img.Name,
		&img.Data,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("No image found for user:", userID)
			return nil, fmt.Errorf("%w: %v", custom_error.ErrImageNotFound, err)
		}
		logger.Error("Failed to get user image", err)
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	return &img, nil
}
