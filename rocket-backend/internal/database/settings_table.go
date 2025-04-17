package database

import (
	"database/sql"
	"fmt"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
)

func (s *service) UpdateSettings(userId uuid.UUID, settings types.SettingsDTO) error {
	query := `UPDATE settings
	          SET profile_image = $1, step_goal = $2
	          WHERE user_id = $3`
	_, err := s.db.Exec(query, settings.ProfImg, settings.StepGoal, userId)
	if err != nil {
		logger.Error("Failed to update settings", err)
		return fmt.Errorf("failed to update settings: %w", err)
	}
	return nil
}

func (s *service) CreateSettings(settings types.Settings) error {
	query := `INSERT INTO settings (id, user_id, profile_image, step_goal)
	          VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, settings.ID, settings.UserId, settings.ProfileImage, settings.StepGoal)
	if err != nil {
		logger.Error("Failed to create settings", err)
		return fmt.Errorf("failed to create settings: %w", err)
	}
	return nil
}

func (s *service) GetSettingsByUserID(userID uuid.UUID) (*types.Settings, error) {
	query := `SELECT id, user_id, profile_image, step_goal
	          FROM settings
	          WHERE user_id = $1`

	var settings types.Settings
	err := s.db.QueryRow(query, userID).Scan(
		&settings.ID,
		&settings.UserId,
		&settings.ProfileImage,
		&settings.StepGoal,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("No settings found for user_id:", userID)
			return nil, nil
		}
		logger.Error("Failed to retrieve settings for user_id:", userID, err)
		return nil, fmt.Errorf("failed to retrieve settings: %w", err)
	}

	return &settings, nil
}
