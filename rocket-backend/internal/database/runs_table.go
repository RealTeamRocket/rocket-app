package database

import (
	"fmt"
	"rocket-backend/internal/custom_error"
	"rocket-backend/internal/types"
	"rocket-backend/pkg/logger"

	"github.com/google/uuid"
	"sort"
)

func (s *service) AddRun(userID uuid.UUID, lineString string) error {
	_, err := s.db.Exec(`
		INSERT INTO runs (user_id, linestring, created_at)
		VALUES ($1, ST_GeomFromText($2, 4326), NOW())
	`, userID, lineString)

	if err != nil {
		logger.Error("Failed to add run", err)
		return fmt.Errorf("%w: failed to add run", custom_error.ErrFailedToSave)
	}

	return nil
}

func (s *service) GetLatestRoute(userID uuid.UUID) (types.Run, error) {
	query := `
		SELECT id, user_id, route, created_at
		FROM runs
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 1
	`

	var run types.Run
	err := s.db.QueryRow(query, userID).Scan(&run.ID, &run.UserID, &run.Route, &run.CreatedAt)
	if err != nil {
		logger.Error("Failed to get latest route", err)
		return types.Run{}, fmt.Errorf("%w: failed to retrieve latest route", custom_error.ErrFailedToRetrieveData)
	}

	return run, nil
}

func (s *service) GetRouteByID(userID uuid.UUID, routeID int) (types.Run, error) {
	query := `
		SELECT id, user_id, route, created_at
		FROM runs
		WHERE user_id = $1 AND id = $2
	`

	var run types.Run
	err := s.db.QueryRow(query, routeID, userID).Scan(&run.ID, &run.UserID, &run.Route, &run.CreatedAt)
	if err != nil {
		logger.Error("Failed to get route by ID", err)
		return types.Run{}, fmt.Errorf("%w: failed to retrieve route by ID", custom_error.ErrFailedToRetrieveData)
	}

	return run, nil
}

func (s *service) DeleteRunByID(userID uuid.UUID, runID int) error {
	_, err := s.db.Exec(`
		DELETE FROM runs
		WHERE user_id = $1 AND id = $2
	`, runID, userID)

	if err != nil {
		logger.Error("Failed to delete run", err)
		return fmt.Errorf("%w: failed to delete run", custom_error.ErrFailedToDelete)
	}

	return nil
}