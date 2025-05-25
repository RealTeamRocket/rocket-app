package database

import (
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
