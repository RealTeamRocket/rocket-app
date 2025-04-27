package database

import (
	"fmt"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
	"sort"
)

func (s *service) AddFriend(userID, friendID uuid.UUID) error {
	if userID == friendID {
		return fmt.Errorf("user cannot add themselves as a friend")
	}

	_, err := s.db.Exec(`
		INSERT INTO friends (user_id, friend_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, userID, friendID)

	if err != nil {
		logger.Error("Failed to add friend", err)
		return fmt.Errorf("failed to add friend: %w", err)
	}

	return nil
}

func (s *service) GetFriends(userID uuid.UUID) ([]types.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.rocketpoints
		FROM friends f
		JOIN users u ON f.friend_id = u.id
		WHERE f.user_id = $1
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		logger.Error("Failed to get friends", err)
		return nil, fmt.Errorf("failed to get friends: %w", err)
	}
	defer rows.Close()

	var friends []types.User
	for rows.Next() {
		var friend types.User
		if err := rows.Scan(&friend.ID, &friend.Username, &friend.Email, &friend.RocketPoints); err != nil {
			logger.Error("Failed to scan friend row", err)
			return nil, fmt.Errorf("failed to scan friend row: %w", err)
		}
		friends = append(friends, friend)
	}

	// Sort friends alphanumerically by username
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].Username < friends[j].Username
	})

	return friends, nil
}

func (s *service) GetFriendsRankedByPoints(userID uuid.UUID) ([]types.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.rocketpoints
		FROM friends f
		JOIN users u ON f.friend_id = u.id
		WHERE f.user_id = $1
		ORDER BY u.rocketpoints DESC
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		logger.Error("Failed to get friends ranked by points", err)
		return nil, fmt.Errorf("failed to get friends ranked by points: %w", err)
	}
	defer rows.Close()

	var friends []types.User
	for rows.Next() {
		var friend types.User
		if err := rows.Scan(&friend.ID, &friend.Username, &friend.Email, &friend.RocketPoints); err != nil {
			logger.Error("Failed to scan friend row", err)
			return nil, fmt.Errorf("failed to scan friend row: %w", err)
		}
		friends = append(friends, friend)
	}

	return friends, nil
}
