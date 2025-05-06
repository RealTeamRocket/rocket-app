package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	ExecuteRawSQL(query string) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// credentails
	SaveCredentials(creds types.Credentials) error
	GetUserByEmail(username string) (types.Credentials, error)
	CheckEmail(email string) error

	// users
	SaveUserProfile(user types.User) error
	GetUserByID(userID uuid.UUID) (types.User, error)
	UpdateRocketPoints(userID uuid.UUID, rocketPoints int) error
	GetUserIDByName(name string) (uuid.UUID, error)
	GetTopUsers(limit int) ([]types.User, error)
	GetAllUsers() ([]types.User, error)

	// daily_steps
	UpdateDailySteps(userID uuid.UUID, steps int) error

	// settings
	GetSettingsByUserID(userID uuid.UUID) (*types.Settings, error)
	CreateSettings(settings types.Settings) error
	UpdateSettings(userId uuid.UUID, settings types.SettingsDTO, imageID uuid.UUID) error

	// images
	SaveImage(filename string, data []byte) (uuid.UUID, error)
	GetUserImage(userID uuid.UUID) (*types.UserImage, error)

	// challenges
	GetAllChallenges() ([]types.Challenge, error)
	AssignChallengesToUser(userID uuid.UUID, challenges []types.Challenge) error
	GetUserDailyChallenges(userID uuid.UUID) ([]types.Challenge, error)
	ResetDailyChallenges() error
	InsertChallenge(challenge types.Challenge) error
	CompleteChallenge(UserID uuid.UUID, dto types.CompleteChallengesDTO) error
	IsNewDayForUser(userID uuid.UUID) (bool, error)
	CleanUpChallengesForUser(userID uuid.UUID) error

	// friends
	AddFriend(userID, friendID uuid.UUID) error
	GetFriends(userID uuid.UUID) ([]types.User, error)
	GetFriendsRankedByPoints(userID uuid.UUID) ([]types.User, error)
	DeleteFriend(userID, friendID uuid.UUID) error
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
	fmt.Printf("Connection String is this: %s \n", connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		logger.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func NewWithConfig(connStr string) Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		logger.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) ExecuteRawSQL(query string) (sql.Result, error) {
	return s.db.Exec(query)
}

func (s *service) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.db.QueryRow(query, args...)
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
		logger.Fatal("db down: %v", err)
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
	logger.Debug("Disconnected from database: %s", database)
	return s.db.Close()
}
