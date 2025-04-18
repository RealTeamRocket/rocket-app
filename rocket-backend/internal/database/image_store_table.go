package database

import "github.com/google/uuid"

func (s *service) SaveImage(filename string, data []byte) (uuid.UUID, error) {
	id := uuid.New()

	_, err := s.db.Exec(`
		INSERT INTO image_store (id, image_name, image_data)
		VALUES ($1, $2, $3)
	`, id, filename, data)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
