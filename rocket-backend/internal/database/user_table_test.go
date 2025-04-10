package database

import (
	"fmt"
	"rocket-backend/internal/types"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type UsersTestContext struct {
	srv *service
}

func (c *UsersTestContext) beforeEach() {
	// Setup: Create an instance of the service, potentially connecting to a test DB
	c.srv = &service{db: testDbInstance}
}

func (c *UsersTestContext) afterEach() {
	// Teardown: Clean up resources after the test (e.g., remove test data, reset states)
	_, err := c.srv.db.Exec("DELETE FROM users")
	if err != nil {
		// Log the error if cleanup fails
		fmt.Println("Failed to cleanup test users:", err)
	}
}

func usersTestCase(test func(t *testing.T, c *UsersTestContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &UsersTestContext{}
		context.beforeEach()
		defer context.afterEach()
		test(t, context)
	}
}

func TestSaveUserProfile(t *testing.T) {
	t.Run("TestCheckEmail", usersTestCase(func(t *testing.T, c *UsersTestContext) {
		id := uuid.New()
		username := "johndoe"
		email := "john@doe.com"
		rocketPoints := 100

		user := types.User{
			ID:           id,
			Username:     username,
			Email:        email,
			RocketPoints: rocketPoints,
		}

		err := c.srv.SaveUserProfile(user)
		assert.NoError(t, err, "expected SaveUserProfile to return no error")

		// Verify that the user profile was saved correctly
		var savedUser types.User
		query := `SELECT id, username, email, rocketpoints FROM users WHERE id = $1`
		err = c.srv.db.QueryRow(query, id).Scan(&savedUser.ID, &savedUser.Username, &savedUser.Email, &savedUser.RocketPoints)
		assert.NoError(t, err, "expected query to return no error")
		assert.Equal(t, user.ID, savedUser.ID, "expected IDs to match")
		assert.Equal(t, user.Username, savedUser.Username, "expected usernames to match")
		assert.Equal(t, user.Email, savedUser.Email, "expected emails to match")
		assert.Equal(t, user.RocketPoints, savedUser.RocketPoints, "expected rocket points to match")
	}))
}

func TestGetUserByEmail(t *testing.T) {
	t.Run("TestCheckEmail", usersTestCase(func(t *testing.T, c *UsersTestContext) {
		id := uuid.New()
		username := "johndoe"
		email := "john@doe.com"
		rocketPoints := 100

		user := types.User{
			ID:           id,
			Username:     username,
			Email:        email,
			RocketPoints: rocketPoints,
		}

		err := c.srv.SaveUserProfile(user)
		assert.NoError(t, err, "expected SaveUserProfile to return no error")

		// Verify that the user profile can be retrieved correctly
		var savedUser types.User
		query := `SELECT id, username, email, rocketpoints FROM users WHERE id = $1`
		err = c.srv.db.QueryRow(query, id).Scan(&savedUser.ID, &savedUser.Username, &savedUser.Email, &savedUser.RocketPoints)
		assert.NoError(t, err, "expected query to return no error")
		assert.Equal(t, user.ID, savedUser.ID, "expected IDs to match")
		assert.Equal(t, user.Username, savedUser.Username, "expected usernames to match")
		assert.Equal(t, user.Email, savedUser.Email, "expected emails to match")
		assert.Equal(t, user.RocketPoints, savedUser.RocketPoints, "expected rocket points to match")
	}))
}
