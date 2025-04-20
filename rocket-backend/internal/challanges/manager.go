package challanges

import (
	"math/rand"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"
	"time"
)

func GetDailies() ([]types.Challenge, error) {
	challenges, err := loadChallenges("internal/challanges/challanges.json")
	if err != nil {
		logger.Error("File not found", err)
		return []types.Challenge{}, err
	}

	if len(challenges) < 5 {
		logger.Warn("Less than 5 challenges available, returning all")
		return []types.Challenge{}, nil
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	rng.Shuffle(len(challenges), func(i, j int) {
		challenges[i], challenges[j] = challenges[j], challenges[i]
	})

	return challenges[:5], nil
}
