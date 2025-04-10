package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	srv := &service{db: testDbInstance}

	stats := srv.Health()

	assert.Equal(t, "up", stats["status"], "expected status to be up")
	assert.NotContains(t, stats, "error", "expected error not to be present")
	assert.Equal(t, "It's healthy", stats["message"], "expected message to be 'It's healthy'")
}
