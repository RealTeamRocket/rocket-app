package server_tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Runs Table Integration", func() {
	var userID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		now := time.Now()
		// Insert into credentials first (to satisfy users FK)
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, "runuser@example.com", "hashedpassword", now, now)
		Expect(err).To(BeNil())

		// Insert into users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "runuser", "runuser@example.com", 0)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM runs")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve a run for a user", func() {
		route := "LINESTRING(0 0,1 1)"
		duration := "45"
		distance := 10.5

		_, err := testDbInstance.Exec(`
			INSERT INTO runs (user_id, route, duration, distance)
			VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
		`, userID, route, duration, distance)
		Expect(err).To(BeNil())

		rows, err := testDbInstance.Query(`
			SELECT id, ST_AsText(route), duration, distance, created_at
			FROM runs
			WHERE user_id = $1
		`, userID)
		Expect(err).To(BeNil())
		defer rows.Close()

		var found bool
		for rows.Next() {
			var id uuid.UUID
			var gotRoute, gotDuration string
			var gotDistance float32
			var createdAt time.Time
			err := rows.Scan(&id, &gotRoute, &gotDuration, &gotDistance, &createdAt)
			Expect(err).To(BeNil())
			Expect(gotRoute).To(Equal(route))
			Expect(gotDuration).To(Equal(duration))
			Expect(gotDistance).To(BeNumerically("~", distance, 0.01))
			found = true
		}
		Expect(found).To(BeTrue())
	})

	It("should delete a run", func() {
		route := "LINESTRING(0 0,1 1)"
		duration := "30"
		distance := 5.0

		var runID uuid.UUID
		err := testDbInstance.QueryRow(`
			INSERT INTO runs (user_id, route, duration, distance)
			VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
			RETURNING id
		`, userID, route, duration, distance).Scan(&runID)
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			DELETE FROM runs WHERE id = $1
		`, runID)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT COUNT(*) FROM runs WHERE id = $1
		`, runID)
		var count int
		err = row.Scan(&count)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(0))
	})

	It("should enforce foreign key constraint on user_id", func() {
		nonExistentUserID := uuid.New()
		route := "LINESTRING(0 0,1 1)"
		duration := "15"
		distance := 2.5

		_, err := testDbInstance.Exec(`
			INSERT INTO runs (user_id, route, duration, distance)
			VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
		`, nonExistentUserID, route, duration, distance)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates foreign key constraint"))
	})
})

var _ = Describe("Planned Runs Table Integration", func() {
	var userID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		now := time.Now()
		// Insert into credentials first (to satisfy users FK)
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, "planuser@example.com", "hashedpassword", now, now)
		Expect(err).To(BeNil())

		// Insert into users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "planuser", "planuser@example.com", 0)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM planned_runs")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve a planned run for a user", func() {
		route := "LINESTRING(2 2,3 3)"
		name := "Morning Plan"
		distance := 5.0

		_, err := testDbInstance.Exec(`
			INSERT INTO planned_runs (user_id, route, name, distance)
			VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
		`, userID, route, name, distance)
		Expect(err).To(BeNil())

		rows, err := testDbInstance.Query(`
			SELECT id, ST_AsText(route), name, created_at, distance
			FROM planned_runs
			WHERE user_id = $1
		`, userID)
		Expect(err).To(BeNil())
		defer rows.Close()

		var found bool
		for rows.Next() {
			var id uuid.UUID
			var gotRoute, gotName string
			var createdAt time.Time
			var gotDistance float64
			err := rows.Scan(&id, &gotRoute, &gotName, &createdAt, &gotDistance)
			Expect(err).To(BeNil())
			Expect(gotRoute).To(Equal(route))
			Expect(gotName).To(Equal(name))
			Expect(gotDistance).To(BeNumerically("~", distance, 0.01))
			found = true
		}
		Expect(found).To(BeTrue())
	})

	It("should delete a planned run", func() {
		route := "LINESTRING(2 2,3 3)"
		name := "Delete Plan"
		distance := 3.0

		var planID uuid.UUID
		err := testDbInstance.QueryRow(`
			INSERT INTO planned_runs (user_id, route, name, distance)
			VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
			RETURNING id
		`, userID, route, name, distance).Scan(&planID)
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			DELETE FROM planned_runs WHERE id = $1
		`, planID)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT COUNT(*) FROM planned_runs WHERE id = $1
		`, planID)
		var count int
		err = row.Scan(&count)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(0))
	})

	It("should enforce unique constraint on (user_id, name)", func() {
		route := "LINESTRING(2 2,3 3)"
		name := "Unique Plan"
		distance := 4.0

		_, err := testDbInstance.Exec(`
			INSERT INTO planned_runs (user_id, route, name, distance)
			VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
		`, userID, route, name, distance)
		Expect(err).To(BeNil())

		// Try to insert another planned run with the same user_id and name
		_, err = testDbInstance.Exec(`
			INSERT INTO planned_runs (user_id, route, name, distance)
			VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
		`, userID, route, name, distance)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("duplicate key"))
	})

	It("should enforce foreign key constraint on user_id", func() {
		nonExistentUserID := uuid.New()
		route := "LINESTRING(2 2,3 3)"
		name := "FK Plan"
		distance := 2.5

		_, err := testDbInstance.Exec(`
			INSERT INTO planned_runs (user_id, route, name, distance)
			VALUES ($1, ST_GeomFromText($2, 4326), $3, $4)
		`, nonExistentUserID, route, name, distance)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates foreign key constraint"))
	})
})
