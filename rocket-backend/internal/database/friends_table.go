package database

import (
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
	"sort"
)

func (s *service) AddFriend(userID, friendID uuid.UUID) error {
	if userID == friendID {
		return custom_error.ErrFailedToSave
	}

	_, err := s.db.Exec(`
		INSERT INTO friends (user_id, friend_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, userID, friendID)

	if err != nil {
		logger.Error("Failed to add friend", err)
		return fmt.Errorf("%w: failed to add friend", custom_error.ErrFailedToSave)
	}

	return nil
}

func (s *service) DeleteFriend(userID, friendID uuid.UUID) error {
	result, err := s.db.Exec(`
		DELETE FROM friends
		WHERE user_id = $1 AND friend_id = $2
	`, userID, friendID)

	if err != nil {
		logger.Error("Failed to delete friend", err)
		return fmt.Errorf("%w: failed to delete friend", custom_error.ErrFailedToDelete)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("Failed to get rows affected", err)
		return fmt.Errorf("%w: failed to delete friend", custom_error.ErrFailedToDelete)
	}

	if rowsAffected == 0 {
		return custom_error.ErrUserNotFound
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
		return nil, fmt.Errorf("%w: failed to retrieve friends", custom_error.ErrFailedToRetrieveData)
	}
	defer rows.Close()

	var friends []types.User
	for rows.Next() {
		var friend types.User
		if err := rows.Scan(&friend.ID, &friend.Username, &friend.Email, &friend.RocketPoints); err != nil {
			logger.Error("Failed to scan friend row", err)
			return nil, fmt.Errorf("%w: failed to scan friend row", custom_error.ErrFailedToRetrieveData)
		}
		friends = append(friends, friend)
	}

	// Sort friends alphanumerically by username
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].Username < friends[j].Username
	})

	if friends == nil {
		friends = []types.User{}
	}

	return friends, nil
}

func (s *service) GetFriendsRankedByPoints(userID uuid.UUID) ([]types.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.rocketpoints
		FROM friends f
		JOIN users u ON f.friend_id = u.id
		WHERE f.user_id = $1
	`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		logger.Error("Failed to get friends", err)
		return nil, fmt.Errorf("%w: failed to retrieve friends", custom_error.ErrFailedToRetrieveData)
	}
	defer rows.Close()

	var friends []types.User
	for rows.Next() {
		var friend types.User
		if err := rows.Scan(&friend.ID, &friend.Username, &friend.Email, &friend.RocketPoints); err != nil {
			logger.Error("Failed to scan friend row", err)
			return nil, fmt.Errorf("%w: failed to scan friend row", custom_error.ErrFailedToRetrieveData)
		}
		friends = append(friends, friend)
	}

	// Also include the user themselves
	var user types.User
	userQuery := `
		SELECT id, username, email, rocketpoints
		FROM users
		WHERE id = $1
	`
	err = s.db.QueryRow(userQuery, userID).Scan(&user.ID, &user.Username, &user.Email, &user.RocketPoints)
	if err != nil {
		logger.Error("Failed to get user for self-inclusion in ranking", err)
	} else {
		friends = append(friends, user)
	}

	// Sort by rocketpoints descending
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].RocketPoints > friends[j].RocketPoints
	})

	if friends == nil {
		friends = []types.User{}
	}

	return friends, nil
}

func (s *service) GetFollowers(userID uuid.UUID) ([]types.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.rocketpoints
		FROM friends f
		JOIN users u ON f.user_id = u.id
		WHERE f.friend_id = $1
	`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		logger.Error("Failed to get followers", err)
		return nil, fmt.Errorf("%w: failed to retrieve followers", custom_error.ErrFailedToRetrieveData)
	}
	defer rows.Close()

	var followers []types.User
	for rows.Next() {
		var follower types.User
		if err := rows.Scan(&follower.ID, &follower.Username, &follower.Email, &follower.RocketPoints); err != nil {
			logger.Error("Failed to scan follower row", err)
			return nil, fmt.Errorf("%w: failed to scan follower row", custom_error.ErrFailedToRetrieveData)
		}
		followers = append(followers, follower)
	}

	// Sort followers alphanumerically by username
	sort.Slice(followers, func(i, j int) bool {
		return followers[i].Username < followers[j].Username
	})

	if followers == nil {
		followers = []types.User{}
	}

	return followers, nil
}
