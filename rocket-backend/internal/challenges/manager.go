package challenges

import (
	"fmt"
	"os"
	"path/filepath"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/database"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
)

type ChallengeManager struct {
	db database.Service
}

func NewChallengeManager(db database.Service) *ChallengeManager {
	return &ChallengeManager{db: db}
}

func (cm *ChallengeManager) GetDailies(userID uuid.UUID) ([]types.Challenge, error) {
	// Ensure challenges are loaded into the database
	err := cm.ensureChallengesLoaded()
	if err != nil {
		logger.Error("Failed to ensure challenges are loaded", err)
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	// Fetch daily challenges for the user
	dailies, err := cm.db.GetUserDailyChallenges(userID)
	if err != nil {
		logger.Error("Failed to fetch user daily challenges", err)
		return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	// If no challenges are assigned yet, assign new ones
	if len(dailies) == 0 {
		allChallenges, err := cm.db.GetAllChallenges()
		if err != nil {
			return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
		}

		if len(allChallenges) < 5 {
			logger.Warn("Less than 5 challenges available, returning all")
			return nil, fmt.Errorf("%w: not enough challenges available: got %d, need at least 5", custom_error.ErrChallengeNotFound, len(allChallenges))
		}

		// Shuffle and pick 5 random challenges
		shuffledChallenges := database.ShuffleChallenges(allChallenges)
		dailyChallenges := shuffledChallenges[:5]

		err = cm.db.AssignChallengesToUser(userID, dailyChallenges)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
		}

		return dailyChallenges, nil
	}

	return dailies, nil
}

func (cm *ChallengeManager) ensureChallengesLoaded() error {
	existingChallenges, err := cm.db.GetAllChallenges()
	if err != nil {
		logger.Error("Failed to fetch challenges from database", err)
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	if len(existingChallenges) > 0 {
		return nil
	}

	wd, err := os.Getwd()
	if err != nil {
		logger.Error("Unable to get working directory", err)
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	filePath := filepath.Join(wd, "internal", "challenges", "challenges.json")
	challenges, err := LoadChallengesFromFile(filePath)
	if err != nil {
		logger.Error("Failed to load challenges from file", err)
		return fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, err)
	}

	for _, challenge := range challenges {
		err := cm.db.InsertChallenge(challenge)
		if err != nil {
			logger.Error("Failed to insert challenge into database", err)
			return fmt.Errorf("%w: %v", custom_error.ErrFailedToSave, err)
		}
	}

	logger.Info("Challenges successfully loaded into the database.")
	return nil
}
