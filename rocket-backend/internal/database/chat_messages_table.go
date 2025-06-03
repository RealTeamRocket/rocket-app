package database

import (
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
)

func (s *service) SaveChatMessage(userID uuid.UUID, message string, timestamp string) error {

	query := `
		INSERT INTO chat_messages (user_id, message, timestamp)
		VALUES ($1, $2, $3)
	`
	_, err := s.db.Exec(query, userID, message, timestamp)
	if err != nil {
		logger.Error("Failed to save chat message", err)
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
	}
	return nil
}
