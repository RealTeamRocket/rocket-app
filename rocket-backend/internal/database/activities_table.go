package database

import (
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"
	"time"

	"github.com/google/uuid"
)

func (s *service) SaveActivity(userID uuid.UUID, message string) error {
	id := uuid.New()
	now := time.Now()

	query := `
		INSERT INTO activities (id, user_id, time, message)
		VALUES ($1, $2, $3, $4)
	`
	_, err := s.db.Exec(query, id, userID, now, message)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to save activity: %v", err))
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
	}
	return nil
}

func (s *service) GetActivitiesForUserAndFriends(userID uuid.UUID) ([]types.ActivityWithUser, error) {
	query := `
		SELECT
			a.user_id,
			CASE WHEN a.user_id = $1 THEN 'You' ELSE u.username END AS name,
			a.time,
			a.message
		FROM activities a
		JOIN users u ON a.user_id = u.id
		WHERE
			(a.user_id = $1
			OR a.user_id IN (
				SELECT friend_id FROM friends WHERE user_id = $1
			))
			AND a.time::date = CURRENT_DATE
		ORDER BY a.time DESC
	`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to fetch activities for user and friends: %v", err))
		return nil, custom_error.ErrFailedToRetrieveData
	}
	defer rows.Close()

	var activities []types.ActivityWithUser
	for rows.Next() {
		var activity types.ActivityWithUser
		if err := rows.Scan(&activity.UserID, &activity.Name, &activity.Time, &activity.Message); err != nil {
			logger.Error(fmt.Sprintf("Failed to scan activity row: %v", err))
			return nil, custom_error.ErrDatabaseQuery
		}
		activities = append(activities, activity)
	}
	if err := rows.Err(); err != nil {
		logger.Error(fmt.Sprintf("Error iterating over activity rows: %v", err))
		return nil, custom_error.ErrDatabaseQuery
	}
	return activities, nil
}
