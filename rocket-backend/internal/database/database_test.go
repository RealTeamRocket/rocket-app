package database

import (
    "context"
    "fmt"
    "log"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
    teardown, err := mustStartPostgresContainer()
    if err != nil {
        log.Fatalf("could not start postgres container: %v", err)
    }

    dbAddr := fmt.Sprintf("%s:%s", host, port)
    err = migrateDb(dbAddr)
    if err != nil {
        log.Fatalf("could not migrate database: %v", err)
    }

    m.Run()

    if teardown != nil && teardown(context.Background()) != nil {
        log.Fatalf("could not teardown postgres container: %v", err)
    }
}

func TestNew(t *testing.T) {
    srv := New()
    assert.NotNil(t, srv, "New() returned nil")
}

func TestHealth(t *testing.T) {
    srv := New()

    stats := srv.Health()

    assert.Equal(t, "up", stats["status"], "expected status to be up")
    assert.NotContains(t, stats, "error", "expected error not to be present")
    assert.Equal(t, "It's healthy", stats["message"], "expected message to be 'It's healthy'")
}

func TestClose(t *testing.T) {
    srv := New()

    assert.NoError(t, srv.Close(), "expected Close() to return nil")
}
