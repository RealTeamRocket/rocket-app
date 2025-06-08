package server_tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Daily Steps Table Integration", func() {
	var userID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		now := time.Now()
		// Insert into credentials first (to satisfy users FK)
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, "stepsuser@example.com", "hashedpassword", now, now)
		Expect(err).To(BeNil())

		// Insert into users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "stepsuser", "stepsuser@example.com", 0)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM daily_steps")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve daily steps for a user", func() {
		today := time.Now().Format("2006-01-02")
		id := uuid.New()
		steps := 1234

		_, err := testDbInstance.Exec(`
			INSERT INTO daily_steps (id, user_id, steps_taken, date)
			VALUES ($1, $2, $3, $4)
		`, id, userID, steps, today)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT steps_taken FROM daily_steps WHERE user_id = $1 AND date = $2
		`, userID, today)
		var gotSteps int
		err = row.Scan(&gotSteps)
		Expect(err).To(BeNil())
		Expect(gotSteps).To(Equal(steps))
	})

	It("should update steps for the same day", func() {
		today := time.Now().Format("2006-01-02")
		id := uuid.New()
		_, err := testDbInstance.Exec(`
			INSERT INTO daily_steps (id, user_id, steps_taken, date)
			VALUES ($1, $2, $3, $4)
		`, id, userID, 1000, today)
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			UPDATE daily_steps SET steps_taken = $1 WHERE user_id = $2 AND date = $3
		`, 2000, userID, today)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT steps_taken FROM daily_steps WHERE user_id = $1 AND date = $2
		`, userID, today)
		var gotSteps int
		err = row.Scan(&gotSteps)
		Expect(err).To(BeNil())
		Expect(gotSteps).To(Equal(2000))
	})

	It("should enforce foreign key constraint on user_id", func() {
		nonExistentUserID := uuid.New()
		id := uuid.New()
		today := time.Now().Format("2006-01-02")
		_, err := testDbInstance.Exec(`
			INSERT INTO daily_steps (id, user_id, steps_taken, date)
			VALUES ($1, $2, $3, $4)
		`, id, nonExistentUserID, 500, today)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates foreign key constraint"))
	})
})
