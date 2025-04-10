package database

import (
	"database/sql"
	"os"
	"testing"

)

var testDbInstance *sql.DB

func TestMain(m *testing.M) {
	testDB := SetupTestDatabase()
	testDbInstance = testDB.DbInstance

	defer testDB.TearDown()

	os.Exit(m.Run())
}
