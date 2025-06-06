package database

import (
	"fmt"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"
	"time"

	"github.com/google/uuid"
)

func (s *service) saveStepMilestoneActivities(userID uuid.UUID, oldSteps, newSteps int) {
	milestone := 2000
	for threshold := ((oldSteps/milestone)+1)*milestone; threshold <= newSteps; threshold += milestone {
		message := fmt.Sprintf("🎉 Has reached %d steps today! 🚀", threshold)
		_ = s.SaveActivity(userID, message)
	}
}

func (s *service) UpdateDailySteps(userID uuid.UUID, steps int) error {
	currentDate := time.Now().Format("2006-01-02")

	var existingSteps int
	queryCheck := `SELECT steps_taken FROM daily_steps WHERE user_id = $1 AND date = $2`
	err := s.db.QueryRow(queryCheck, userID, currentDate).Scan(&existingSteps)

	if err != nil {
		id := uuid.New()
		queryInsert := `INSERT INTO daily_steps (id, user_id, steps_taken, date) VALUES ($1, $2, $3, $4)`
		_, err := s.db.Exec(queryInsert, id, userID, steps, currentDate)
		if err != nil {
			logger.Error("Error inserting daily steps: %v\n", err)
			return fmt.Errorf("failed to insert daily steps: %w", err)
		}
		rocketPoints := steps / 10
		if rocketPoints > 0 {
			if err := s.UpdateRocketPoints(userID, rocketPoints); err != nil {
				logger.Error("Error updating rocket points: %v\n", err)
				return fmt.Errorf("failed to update rocket points: %w", err)
			}
		}
		s.saveStepMilestoneActivities(userID, 0, steps)
	} else {
		// If an entry exists, update the steps and rocket points only if the new steps are greater than the existing steps
		if steps > existingSteps {
			stepDiff := steps - existingSteps
			if stepDiff > 0 {
				queryUpdate := `UPDATE daily_steps SET steps_taken = $1 WHERE user_id = $2 AND date = $3`
				_, err := s.db.Exec(queryUpdate, steps, userID, currentDate)
				if err != nil {
					logger.Error("Error updating daily steps: %v\n", err)
					return fmt.Errorf("failed to update daily steps: %w", err)
				}
				rocketPoints := stepDiff / 10
				if rocketPoints > 0 {
					if err := s.UpdateRocketPoints(userID, rocketPoints); err != nil {
						logger.Error("Error updating rocket points: %v\n", err)
						return fmt.Errorf("failed to update rocket points: %w", err)
					}
				}
				s.saveStepMilestoneActivities(userID, existingSteps, steps)
			}
		}
	}

	return nil
}

func (s *service) GetUserStatistics(userID uuid.UUID) ([]types.StepStatistic, error) {
	query := `
		SELECT date, steps_taken
		FROM daily_steps
		WHERE user_id = $1 AND date >= CURRENT_DATE - INTERVAL '6 days'
		ORDER BY date ASC
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query daily steps: %w", err)
	}
	defer rows.Close()

	// Map to store steps by date
	stepsByDate := make(map[string]int)

	// Parse the results
	for rows.Next() {
		var date time.Time
		var steps int
		if err := rows.Scan(&date, &steps); err != nil {
			return nil, fmt.Errorf("failed to scan daily steps: %w", err)
		}
		stepsByDate[date.Format("2006-01-02")] = steps
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	// Prepare the result slice
	var statistics []types.StepStatistic
	dayNames := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}

	// Iterate over the last 7 days
	for i := 6; i >= 0; i-- {
		date := time.Now().AddDate(0, 0, -i)
		day := dayNames[date.Weekday()]
		dateStr := date.Format("2006-01-02")

		// Get steps for the day, default to 0 if not found
		steps := stepsByDate[dateStr]

		statistics = append(statistics, types.StepStatistic{
			Day:   day,
			Steps: steps,
		})
	}

	return statistics, nil
}
