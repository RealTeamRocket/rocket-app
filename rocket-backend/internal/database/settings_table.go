package database

import (
	"database/sql"
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
)

func (s *service) UpdateSettings(userId uuid.UUID, settings types.SettingsDTO, imageID uuid.UUID) error {
	query := `UPDATE settings
	          SET image_id = $1, step_goal = $2
	          WHERE user_id = $3`
	_, err := s.db.Exec(query, imageID, settings.StepGoal, userId)
	if err != nil {
		logger.Error("Failed to update settings", err)
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToUpdate, err)
	}
	return nil
}

func (s *service) CreateSettings(settings types.Settings) error {
	query := `INSERT INTO settings (id, user_id, image_id, step_goal)
	          VALUES ($1, $2, $3, $4)`

	var imageId interface{}
	if settings.ImageId == uuid.Nil {
		imageId = nil
	} else {
		imageId = settings.ImageId
	}

	_, err := s.db.Exec(query, settings.ID, settings.UserId, imageId, settings.StepGoal)
	if err != nil {
		logger.Error("Failed to create settings", err)
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
	}
	return nil
}

func (s *service) GetSettingsByUserID(userID uuid.UUID) (*types.Settings, error) {
	query := `SELECT id, user_id, image_id, step_goal
	          FROM settings
	          WHERE user_id = $1`

	var settings types.Settings
	err := s.db.QueryRow(query, userID).Scan(
		&settings.ID,
		&settings.UserId,
		&settings.ImageId,
		&settings.StepGoal,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("No settings found for user_id:", userID)
			return nil, fmt.Errorf("%w: %v", custom_error.ErrSettingsNotFound, err)
		}
		logger.Error("Failed to retrieve settings for user_id:", userID, err)
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	return &settings, nil
}
