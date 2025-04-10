package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type dailyStepsTestContext struct {
	srv *service
}

func (c *dailyStepsTestContext) beforeEach() {
	// Setup: Create an instance of the service, potentially connecting to a test DB
	c.srv = &service{db: testDbInstance}
}

func (c *dailyStepsTestContext) afterEach() {
	// Teardown: Clean up resources after the test (e.g., remove test data, reset states)
	_, err := c.srv.db.Exec("DELETE FROM daily_steps") // Ensure you're cleaning up the right table
	if err != nil {
		// Log the error if cleanup fails
		fmt.Println("Failed to cleanup test daily_steps:", err)
	}
}

func dailyStepsTestCase(test func(t *testing.T, c *dailyStepsTestContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &dailyStepsTestContext{}
		context.beforeEach()
		defer context.afterEach()
		test(t, context)
	}
}

// Example test: TestInsertDailySteps
func TestInsertDailySteps(t *testing.T) {
	t.Run("TestInsertDailySteps", dailyStepsTestCase(func(t *testing.T, c *dailyStepsTestContext) {
		userID := uuid.New()
		steps := 10000
		currentDate := time.Now().Format("2006-01-02")

		// Insert daily steps
		err := c.srv.UpdateDailySteps(userID, steps)

		// Verify that the steps were inserted correctly
		var savedSteps int
		err = c.srv.db.QueryRow("SELECT steps_taken FROM daily_steps WHERE user_id = $1 AND date = $2", userID, currentDate).Scan(&savedSteps)
		assert.NoError(t, err, "expected to retrieve steps_taken")
		assert.Equal(t, steps, savedSteps, "expected steps to match")
	}))
}
