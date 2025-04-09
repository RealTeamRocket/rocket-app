package database

import (
    "context"
    "fmt"
    "log"
    "path/filepath"
    "runtime"
    "time"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/modules/postgres"
    "github.com/testcontainers/testcontainers-go/wait"
)

func mustStartPostgresContainer() (func(context.Context, ...testcontainers.TerminateOption) error, error) {
    var (
        dbName = "database"
        dbPwd  = "password"
        dbUser = "user"
    )

    dbContainer, err := postgres.Run(
        context.Background(),
        "postgis/postgis:latest",
        postgres.WithDatabase(dbName),
        postgres.WithUsername(dbUser),
        postgres.WithPassword(dbPwd),
        testcontainers.WithWaitStrategy(
            wait.ForLog("database system is ready to accept connections").
                WithOccurrence(2).
                WithStartupTimeout(5*time.Second)),
    )
    if err != nil {
        return nil, err
    }

    database = dbName
    password = dbPwd
    username = dbUser

    dbHost, err := dbContainer.Host(context.Background())
    if err != nil {
        return dbContainer.Terminate, err
    }

    dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
    if err != nil {
        return dbContainer.Terminate, err
    }

    host = dbHost
    port = dbPort.Port()

    return dbContainer.Terminate, err
}

func migrateDb(dbAddr string) error {
    _, path, _, ok := runtime.Caller(0)
    if !ok {
        return fmt.Errorf("failed to get path")
    }
    pathToMigrationFiles := filepath.Join(filepath.Dir(path), "../../migrations")

    databaseURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, dbAddr, database)
    m, err := migrate.New(fmt.Sprintf("file://%s", pathToMigrationFiles), databaseURL)
    if err != nil {
        return err
    }
    defer m.Close()

    err = m.Up()
    if err != nil && err != migrate.ErrNoChange {
        return err
    }

    log.Println("migration done")

    return nil
}
