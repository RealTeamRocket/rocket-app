package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"rocket-backend/internal/types"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	SaveCredentials(creds types.Credentials) error
	GetUserByEmail(username string) (types.Credentials, error)
	GetUserByID(id uuid.UUID) (types.Credentials, error)
	CheckEmail(email string) error
}

type service struct {
	db *sql.DB
}

var (
	database   = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	schema     = os.Getenv("BLUEPRINT_DB_SCHEMA")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}

func (s *service) SaveCredentials(creds types.Credentials) error {
	query := `INSERT INTO credentials (id, username, email, password, created_at, last_login) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Exec(query, creds.ID, creds.Username, creds.Email, creds.Password, creds.CreatedAt, creds.LastLogin)
	if err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}
	return nil
}

func (s *service) GetUserByEmail(email string) (types.Credentials, error) {
	var creds types.Credentials
	query := `SELECT id, username, email, password, created_at, last_login FROM credentials WHERE email = $1`
	err := s.db.QueryRow(query, email).Scan(&creds.ID, &creds.Username, &creds.Email, &creds.Password, &creds.CreatedAt, &creds.LastLogin)
	if err != nil {
		return creds, fmt.Errorf("failed to get user by username: %w", err)
	}
	return creds, nil
}

func (s *service) GetUserByID(id uuid.UUID) (types.Credentials, error) {
	var creds types.Credentials
	query := `SELECT id, username, email, password, created_at, last_login FROM credentials WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&creds.ID, &creds.Username, &creds.Email, &creds.Password, &creds.CreatedAt, &creds.LastLogin)
	if err != nil {
		return creds, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return creds, nil
}

func (s *service) CheckEmail(email string) error {
	query := `SELECT COUNT(*) FROM credentials WHERE email = $1`
	var count int
	err := s.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check email: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("email already exists")
	}
	return nil
}
