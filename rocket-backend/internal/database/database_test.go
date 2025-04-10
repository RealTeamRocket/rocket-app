package database

import (
	"database/sql"
	"os"
	"rocket-backend/internal/types"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var testDbInstance *sql.DB

func TestMain(m *testing.M) {
	testDB := SetupTestDatabase()
	testDbInstance = testDB.DbInstance

	defer testDB.TearDown()

	os.Exit(m.Run())
}

func TestHealth(t *testing.T) {
	srv := &service{db: testDbInstance}

	stats := srv.Health()

	assert.Equal(t, "up", stats["status"], "expected status to be up")
	assert.NotContains(t, stats, "error", "expected error not to be present")
	assert.Equal(t, "It's healthy", stats["message"], "expected message to be 'It's healthy'")
}

func TestSaveCredential(t *testing.T) {
	srv := &service{db: testDbInstance}

	id := uuid.New()
	email := "john@doe.com"
	password := "securepassword"
	createdAt :=time.Now().Format(time.RFC3339)
	lastLogin :=time.Now().Format(time.RFC3339)

	credentials := types.Credentials{
		ID:        id,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		LastLogin: lastLogin,
	}

	err := srv.SaveCredentials(credentials)
	assert.NoError(t, err, "expected SaveCredentials to return no error")

	// Verify that the credentials were saved correctly
	savedCreds, err := srv.GetUserByEmail(email)
	assert.NoError(t, err, "expected GetUserByEmail to return no error")
	assert.Equal(t, credentials.ID, savedCreds.ID, "expected IDs to match")
	assert.Equal(t, credentials.Email, savedCreds.Email, "expected emails to match")
	assert.Equal(t, credentials.Password, savedCreds.Password, "expected passwords to match")
	assert.Equal(t, credentials.CreatedAt, savedCreds.CreatedAt, "expected created_at to match")
	assert.Equal(t, credentials.LastLogin, savedCreds.LastLogin, "expected last_login to match")
}

func TestCheckEmail (t *testing.T) {
	srv := &service{db: testDbInstance}

	id := uuid.New()
	email := "john@doe.com"
	password := "securepassword"
	createdAt :=time.Now().Format(time.RFC3339)
	lastLogin :=time.Now().Format(time.RFC3339)

	credentials := types.Credentials{
		ID:        id,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		LastLogin: lastLogin,
	}

	err := srv.SaveCredentials(credentials)
	assert.NoError(t, err, "expected SaveCredentials to return no error")

	err = srv.CheckEmail(email)
	assert.Error(t, err, "expected CheckEmail to return an error")
}
