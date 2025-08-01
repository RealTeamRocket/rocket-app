package challenges_tests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"rocket-backend/pkg/logger"
	"runtime"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testDbInstance *sql.DB
var connectionString string

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

const (
	DbName = "test_db"
	DbUser = "test_user"
	DbPass = "test_password"
)

type TestDatabase struct {
	DbInstance *sql.DB
	DbAddress  string
	container  testcontainers.Container
}

var testDB *TestDatabase

var _ = BeforeSuite(func() {
	testDB = SetupTestDatabase()
	testDbInstance = testDB.DbInstance

})

var _ = AfterSuite(func() {
	testDB.TearDown()
})

var _ = AfterEach(func() {
	err := truncateTables(testDbInstance)
	Expect(err).To(BeNil())
})

func SetupTestDatabase() *TestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	container, dbInstance, dbAddr, err := createContainer(ctx)
	if err != nil {
		logger.Fatal("failed to setup test", err)
	}

	err = migrateDb(dbAddr)
	if err != nil {
		logger.Fatal("failed to perform db migration", err)
	}

	return &TestDatabase{
		container:  container,
		DbInstance: dbInstance,
		DbAddress:  dbAddr,
	}
}

func (tdb *TestDatabase) TearDown() {
	tdb.DbInstance.Close()
	_ = tdb.container.Terminate(context.Background())
}

func createContainer(ctx context.Context) (testcontainers.Container, *sql.DB, string, error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": DbPass,
		"POSTGRES_USER":     DbUser,
		"POSTGRES_DB":       DbName,
	}
	var port = "5432/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgis/postgis:12-3.0",
			ExposedPorts: []string{port},
			Env:          env,
			WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, nil, "", fmt.Errorf("failed to start container: %w", err)
	}

	p, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to get container external port: %w", err)
	}

	dbAddr := fmt.Sprintf("localhost:%s", p.Port())
	connectionString = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DbUser, DbPass, dbAddr, DbName)

	// Attempt to open DB connection
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return container, nil, dbAddr, fmt.Errorf("failed to open db connection: %w", err)
	}

	// Wait for DB to be ready by pinging it
	maxWait := time.Second * 10
	start := time.Now()
	for {
		if time.Since(start) > maxWait {
			return container, nil, dbAddr, fmt.Errorf("timeout waiting for database to be ready")
		}

		if err := db.Ping(); err == nil {
			break
		}

		logger.Debug("waiting for db to be ready...")
		time.Sleep(time.Second)
	}

	logger.Info("postgres container ready and running at port:", p.Port())

	return container, db, dbAddr, nil
}

func migrateDb(dbAddr string) error {
	// get location of test
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get path")
	}
	pathToMigrationFiles := filepath.Join(filepath.Dir(path), "../../../migrations")

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DbUser, DbPass, dbAddr, DbName)
	m, err := migrate.New(fmt.Sprintf("file:%s", pathToMigrationFiles), databaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	logger.Debug("migration done")

	return nil
}

// Truncate all tables for test isolation
func truncateTables(db *sql.DB) error {
	_, err := db.Exec(`
		DO $$
		DECLARE
			r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
				EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' RESTART IDENTITY CASCADE;';
			END LOOP;
		END$$;
	`)
	return err
}

func TestChallengesIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Database Integration Tests Suite")
}
