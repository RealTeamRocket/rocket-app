package challenges

import (
	"encoding/json"
	"os"
	"sync"

	"rocket-backend/internal/types"
)

var (
	once        sync.Once
	cached      []types.Challenge
	loadErr     error
)

func LoadChallengesFromFile(path string) ([]types.Challenge, error) {
	once.Do(func() {
		var file *os.File
		file, loadErr = os.Open(path)
		if loadErr != nil {
			return
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		loadErr = decoder.Decode(&cached)
	})

	return cached, loadErr
}
