package challenges

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
)

var (
	once    sync.Once
	cached  []types.Challenge
	loadErr error
)

func LoadChallengesFromFile(path string) ([]types.Challenge, error) {
	once.Do(func() {
		var file *os.File
		file, loadErr = os.Open(path)
		if loadErr != nil {
			loadErr = fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, loadErr)
			return
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		loadErr = decoder.Decode(&cached)
		if loadErr != nil {
			loadErr = fmt.Errorf("%w: %v", custom_error.ErrFailedToRetrieveData, loadErr)
		}
	})

	return cached, loadErr
}
