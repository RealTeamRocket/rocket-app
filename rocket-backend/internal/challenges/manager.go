package challenges

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"
	"time"
)

func GetDailies() ([]types.Challenge, error) {
	wd, err := os.Getwd()
	if err != nil {
		logger.Error("Unable to get working directory", err)
		return nil, err
	}

	filePath := filepath.Join(wd, "internal", "challenges", "challenges.json")
	challenges, err := loadChallenges(filePath)
	if err != nil {
		logger.Error("File not found", err)
		return nil, err
	}

	if len(challenges) < 5 {
		logger.Warn("Less than 5 challenges available, returning all")
		return nil, fmt.Errorf("not enough challenges available: got %d, need at least 5", len(challenges))
	}

	// Shuffle logic
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	rng.Shuffle(len(challenges), func(i, j int) {
		challenges[i], challenges[j] = challenges[j], challenges[i]
	})

	return challenges[:5], nil
}
