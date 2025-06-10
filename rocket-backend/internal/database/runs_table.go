package database

import (
	"rocket-backend/internal/types"

	"github.com/google/uuid"
)

func (s *service) SaveRun(userID uuid.UUID, route string, duration string, distance float64) error {
    query := `
        INSERT INTO runs (user_id, route, duration, distance)
        VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
    `
    _, err := s.db.Exec(query, userID, route, duration, distance)
    if err != nil {
        return err
    }
    return nil
}

func (s *service) GetAllRunsByUser(userID uuid.UUID) ([]types.RunDTO, error) {
    query := `
        SELECT id, ST_AsText(route), duration, distance, created_at
        FROM runs
        WHERE user_id = $1
        ORDER BY created_at DESC
    `
    rows, err := s.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var runs []types.RunDTO
    for rows.Next() {
        var run types.RunDTO
        if err := rows.Scan(&run.ID, &run.Route, &run.Duration, &run.Distance, &run.CreatedAt); err != nil {
            return nil, err
        }
        runs = append(runs, run)
    }
    return runs, nil
}

func (s *service) DeleteRun(runID uuid.UUID) error {
	query := `
		DELETE FROM runs
		WHERE id = $1
	`
	_, err := s.db.Exec(query, runID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) SavePlannedRun(userID uuid.UUID, route string, name string, distance float64) error {
    query := `
        INSERT INTO planned_runs (user_id, route, name, distance)
        VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
    `
    _, err := s.db.Exec(query, userID, route, name, distance)
    return err
}

func (s *service) GetAllPlannedRunsByUser(userID uuid.UUID) ([]types.PlannedRunDTO, error) {
    query := `
        SELECT id, ST_AsText(route), name, created_at, distance
        FROM planned_runs
        WHERE user_id = $1
        ORDER BY created_at DESC
    `
    rows, err := s.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var runs []types.PlannedRunDTO
    for rows.Next() {
        var run types.PlannedRunDTO
        if err := rows.Scan(&run.ID, &run.Route, &run.Name, &run.CreatedAt, &run.Distance); err != nil {
            return nil, err
        }
        runs = append(runs, run)
    }
    return runs, nil
}
