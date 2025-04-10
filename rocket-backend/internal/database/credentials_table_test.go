package database

import (
	"fmt"
	"rocket-backend/internal/types"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type CredentialsTestContext struct {
	srv *service
}

func (c *CredentialsTestContext) beforeEach() {
	// Setup: Create an instance of the service, potentially connecting to a test DB
	c.srv = &service{db: testDbInstance}
}

func (c *CredentialsTestContext) afterEach() {
	// Teardown: Clean up resources after the test (e.g., remove test data, reset states)
	// Example: Delete test credentials if needed
	_, err := c.srv.db.Exec("DELETE FROM credentials")
	if err != nil {
		// Log the error if cleanup fails
		fmt.Println("Failed to cleanup test credentials:", err)
	}
}

func credentialsTestCase(test func(t *testing.T, c *CredentialsTestContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &CredentialsTestContext{}
		context.beforeEach()
		defer context.afterEach()
		test(t, context)
	}
}

func TestSaveCredential(t *testing.T) {
	t.Run("TestSaveCredential", credentialsTestCase(func(t *testing.T, c *CredentialsTestContext) {
		id := uuid.New()
		email := "john@doe.com"
		password := "securepassword"
		createdAt := time.Now().Format(time.RFC3339)
		lastLogin := time.Now().Format(time.RFC3339)

		credentials := types.Credentials{
			ID:        id,
			Email:     email,
			Password:  password,
			CreatedAt: createdAt,
			LastLogin: lastLogin,
		}

		err := c.srv.SaveCredentials(credentials)
		assert.NoError(t, err, "expected SaveCredentials to return no error")

		// Verify that the credentials were saved correctly
		savedCreds, err := c.srv.GetUserByEmail(email)
		assert.NoError(t, err, "expected GetUserByEmail to return no error")
		assert.Equal(t, credentials.ID, savedCreds.ID, "expected IDs to match")
		assert.Equal(t, credentials.Email, savedCreds.Email, "expected emails to match")
		assert.Equal(t, credentials.Password, savedCreds.Password, "expected passwords to match")
		assert.Equal(t, credentials.CreatedAt, savedCreds.CreatedAt, "expected created_at to match")
		assert.Equal(t, credentials.LastLogin, savedCreds.LastLogin, "expected last_login to match")
	}))
}

func TestCheckEmail(t *testing.T) {
	t.Run("TestCheckEmail", credentialsTestCase(func(t *testing.T, c *CredentialsTestContext) {
		id := uuid.New()
		email := "john@doe.com"
		password := "securepassword"
		createdAt := time.Now().Format(time.RFC3339)
		lastLogin := time.Now().Format(time.RFC3339)

		credentials := types.Credentials{
			ID:        id,
			Email:     email,
			Password:  password,
			CreatedAt: createdAt,
			LastLogin: lastLogin,
		}

		err := c.srv.SaveCredentials(credentials)
		assert.NoError(t, err, "expected SaveCredentials to return no error")

		err = c.srv.CheckEmail(email)
		assert.Error(t, err, "expected CheckEmail to return an error")
	}))
}
