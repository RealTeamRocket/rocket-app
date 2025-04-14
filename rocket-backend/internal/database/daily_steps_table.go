package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (s *service) UpdateDailySteps(userID uuid.UUID, steps int) error {
	currentDate := time.Now().Format("2006-01-02")

	// Check if the user already has a record for today
	var existingSteps int
	queryCheck := `SELECT steps_taken FROM daily_steps WHERE user_id = $1 AND date = $2`
	err := s.db.QueryRow(queryCheck, userID, currentDate).Scan(&existingSteps)

	if err != nil {
		id := uuid.New()
		queryInsert := `INSERT INTO daily_steps (id, user_id, steps_taken, date) VALUES ($1, $2, $3, $4)`
		_, err := s.db.Exec(queryInsert, id, userID, steps, currentDate)
		if err != nil {
			// Log and return error if the insert fails
			fmt.Printf("Error inserting daily steps: %v\n", err)
			return fmt.Errorf("failed to insert daily steps: %w", err)
		}
	} else {
		// If an entry exists, update the steps only if the new steps are greater than the existing steps
		if steps > existingSteps {
			queryUpdate := `UPDATE daily_steps SET steps_taken = $1 WHERE user_id = $2 AND date = $3`
			_, err := s.db.Exec(queryUpdate, steps, userID, currentDate)
			if err != nil {
				// Log and return error if the update fails
				fmt.Printf("Error updating daily steps: %v\n", err)
				return fmt.Errorf("failed to update daily steps: %w", err)
			}
		}
	}

	return nil
}
