package server_tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Settings Table Integration", func() {
	var userID uuid.UUID
	var settingsID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		settingsID = uuid.New()
		now := time.Now()
		// Insert into credentials first (to satisfy users FK)
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, "settingsuser@example.com", "hashedpassword", now, now)
		Expect(err).To(BeNil())

		// Insert into users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "settingsuser", "settingsuser@example.com", 0)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM settings")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve settings for a user", func() {
		stepGoal := 12345
		_, err := testDbInstance.Exec(`
			INSERT INTO settings (id, user_id, step_goal)
			VALUES ($1, $2, $3)
		`, settingsID, userID, stepGoal)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT id, user_id, step_goal FROM settings WHERE user_id = $1
		`, userID)
		var gotID uuid.UUID
		var gotUserID uuid.UUID
		var gotStepGoal int
		err = row.Scan(&gotID, &gotUserID, &gotStepGoal)
		Expect(err).To(BeNil())
		Expect(gotID).To(Equal(settingsID))
		Expect(gotUserID).To(Equal(userID))
		Expect(gotStepGoal).To(Equal(stepGoal))
	})

	It("should return no rows if settings do not exist for user", func() {
		otherUserID := uuid.New()
		row := testDbInstance.QueryRow(`
			SELECT id, user_id, step_goal FROM settings WHERE user_id = $1
		`, otherUserID)
		var gotID uuid.UUID
		var gotUserID uuid.UUID
		var gotStepGoal int
		err := row.Scan(&gotID, &gotUserID, &gotStepGoal)
		Expect(err).ToNot(BeNil())
	})
})
