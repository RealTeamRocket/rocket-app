package database

import (
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
)

func (s *service) SaveChatMessage(userID uuid.UUID, message string, timestamp string) (uuid.UUID, error) {
	query := `
		INSERT INTO chat_messages (user_id, message, timestamp)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id uuid.UUID
	err := s.db.QueryRow(query, userID, message, timestamp).Scan(&id)
	if err != nil {
		logger.Error("Failed to save chat message and return ID", err)
		return uuid.Nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
	}
	return id, nil
}

func (s *service) GetChatMessages(userID uuid.UUID) ([]types.ChatMessage, error) {
	query := `
		SELECT cm.id, cm.user_id, u.username, cm.message, cm.timestamp
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
		if err := rows.Scan(&msg.ID, &dbUserID, &msg.Username, &msg.Message, &msg.Timestamp); err != nil {
			logger.Error("Failed to scan chat message", err)
			continue
		}
		if dbUserID == userID {
			msg.Username = "You"
		}
		// Count reactions for this message
		count, err := s.CountReactionsForMessage(msg.ID)
		if err != nil {
			logger.Error("Failed to count reactions for message", err)
			msg.Reactions = 0
		} else {
			msg.Reactions = count
		}
		// Check if the current user has reacted to this message
		hasReacted := false
		checkQuery := `SELECT 1 FROM chat_messages_reactions WHERE message_id = $1 AND user_id = $2 LIMIT 1`
		row := s.db.QueryRow(checkQuery, msg.ID, userID)
		var dummy int
		if err := row.Scan(&dummy); err == nil {
			hasReacted = true
		}
		msg.HasReacted = hasReacted
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		logger.Error("Error iterating chat messages", err)
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToLoad, err)
	}
	return messages, nil
}

func (s *service) AddReactionToChatMessage(userID uuid.UUID, messageID uuid.UUID) error {
	query := `
        INSERT INTO chat_messages_reactions (message_id, user_id)
        VALUES ($1, $2)
        ON CONFLICT DO NOTHING
    `
    _, err := s.db.Exec(query, messageID, userID)
    return err
}

func (s *service) CountReactionsForMessage(messageID uuid.UUID) (int, error) {
    var count int
    query := `SELECT COUNT(*) FROM chat_messages_reactions WHERE message_id = $1`
    err := s.db.QueryRow(query, messageID).Scan(&count)
    return count, err
}

func (s *service) GetIDByMessageID(messageID uuid.UUID) (uuid.UUID, error) {
	query := `SELECT user_id FROM chat_messages WHERE id = $1`
	var userID uuid.UUID
	err := s.db.QueryRow(query, messageID).Scan(&userID)
	if err != nil {
		logger.Error("Failed to get user ID by message ID", err)
		return uuid.Nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}
	return userID, nil
}
