package database

import (
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
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

func (s *service) GetChatMessages(userID uuid.UUID) ([]types.ChatMessage, error) {
	query := `
		SELECT cm.user_id, u.username, cm.message, cm.timestamp
		FROM chat_messages cm
		JOIN users u ON cm.user_id = u.id
		ORDER BY cm.timestamp ASC
	`
	rows, err := s.db.Query(query)
	if err != nil {
		logger.Error("Failed to load chat messages", err)
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToLoad, err)
	}
	defer rows.Close()

	var messages []types.ChatMessage
	for rows.Next() {
		var msg types.ChatMessage
		var dbUserID uuid.UUID
		if err := rows.Scan(&dbUserID, &msg.Username, &msg.Message, &msg.Timestamp); err != nil {
			logger.Error("Failed to scan chat message", err)
			continue
		}
		if dbUserID == userID {
			msg.Username = "You"
		}
		msg.Reactions = 0
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		logger.Error("Error iterating chat messages", err)
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToLoad, err)
	}
	return messages, nil
}
