package database

import (
	"fmt"
	"math/rand"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"
	"time"

	"github.com/google/uuid"
)

func (s *service) GetAllChallenges() ([]types.Challenge, error) {
	var challenges []types.Challenge
	query := `SELECT id, description AS text, points_reward AS points FROM challenges`
	rows, err := s.db.Query(query)
	if err != nil {
		logger.Error("Failed to fetch challenges from database", err)
		return nil, fmt.Errorf("failed to fetch challenges: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var challenge types.Challenge
		if err := rows.Scan(&challenge.ID, &challenge.Text, &challenge.Points); err != nil {
			logger.Error("Failed to scan challenge row", err)
			return nil, fmt.Errorf("failed to scan challenge row: %w", err)
		}
		challenges = append(challenges, challenge)
	}

	if err := rows.Err(); err != nil {
		logger.Error("Error iterating over challenge rows", err)
		return nil, fmt.Errorf("error iterating over challenge rows: %w", err)
	}

	return challenges, nil
}

func (s *service) AssignChallengesToUser(userID uuid.UUID, challenges []types.Challenge) error {
	query := `
		INSERT INTO user_challenges (user_id, challenge_id, date)
		VALUES ($1, $2, CURRENT_DATE)
		ON CONFLICT (user_id, challenge_id, date) DO NOTHING
	`

	for _, challenge := range challenges {
		_, err := s.db.Exec(query, userID, challenge.ID)
		if err != nil {
			logger.Error("Failed to assign challenge to user", err)
			return fmt.Errorf("failed to assign challenge to user: %w", err)
		}
	}

	return nil
}

func (s *service) GetUserDailyChallenges(userID uuid.UUID) ([]types.Challenge, error) {
	var challenges []types.Challenge
	query := `
		SELECT c.id, c.description AS text, c.points_reward AS points
		FROM user_challenges uc
		JOIN challenges c ON uc.challenge_id = c.id
		WHERE uc.user_id = $1 AND uc.date = CURRENT_DATE
	`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		logger.Error("Failed to fetch user daily challenges", err)
		return nil, fmt.Errorf("failed to fetch user daily challenges: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var challenge types.Challenge
		if err := rows.Scan(&challenge.ID, &challenge.Text, &challenge.Points); err != nil {
			logger.Error("Failed to scan user challenge row", err)
			return nil, fmt.Errorf("failed to scan user challenge row: %w", err)
		}
		challenges = append(challenges, challenge)
	}

	if err := rows.Err(); err != nil {
		logger.Error("Error iterating over user challenge rows", err)
		return nil, fmt.Errorf("error iterating over user challenge rows: %w", err)
	}

	return challenges, nil
}

func (s *service) ResetDailyChallenges() error {
	_, err := s.db.Exec(`DELETE FROM user_challenges WHERE date < CURRENT_DATE`)
	if err != nil {
		logger.Error("Failed to delete old challenges", err)
		return fmt.Errorf("failed to delete old challenges: %w", err)
	}

	// Fetch all users
	var userIDs []uuid.UUID
	query := `SELECT id FROM users`
	rows, err := s.db.Query(query)
	if err != nil {
		logger.Error("Failed to fetch user IDs", err)
		return fmt.Errorf("failed to fetch user IDs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			logger.Error("Failed to scan user ID", err)
			return fmt.Errorf("failed to scan user ID: %w", err)
		}
		userIDs = append(userIDs, userID)
	}

	if err := rows.Err(); err != nil {
		logger.Error("Error iterating over user IDs", err)
		return fmt.Errorf("error iterating over user IDs: %w", err)
	}

	// Assign new challenges to each user
	allChallenges, err := s.GetAllChallenges()
	if err != nil {
		return err
	}

	for _, userID := range userIDs {
		shuffledChallenges := ShuffleChallenges(allChallenges)
		dailyChallenges := shuffledChallenges[:5]

		err := s.AssignChallengesToUser(userID, dailyChallenges)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to assign challenges to user %s", userID), err)
			return err
		}
	}

	return nil
}

func ShuffleChallenges(challenges []types.Challenge) []types.Challenge {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	shuffled := make([]types.Challenge, len(challenges))
	copy(shuffled, challenges)
	rng.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

func (s *service) InsertChallenge(challenge types.Challenge) error {
	query := `
		INSERT INTO challenges (id, description, points_reward)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO NOTHING
	`

	_, err := s.db.Exec(query, challenge.ID, challenge.Text, challenge.Points)
	if err != nil {
		logger.Error("Failed to insert challenge into database", err)
		return fmt.Errorf("failed to insert challenge: %w", err)
	}

	return nil
}

func (s *service) CompleteChallenge(UserID uuid.UUID, dto types.CompleteChallengesDTO) error {
	query := `
		UPDATE user_challenges
		SET is_completed = TRUE
		WHERE user_id = $1 AND challenge_id = $2 AND date = CURRENT_DATE
	`

	_, err := s.db.Exec(query, UserID, dto.ChallengeID)
	if err != nil {
		logger.Error("Failed to mark challenge as completed", err)
		return fmt.Errorf("failed to mark challenge as completed: %w", err)
	}

	return nil
}
