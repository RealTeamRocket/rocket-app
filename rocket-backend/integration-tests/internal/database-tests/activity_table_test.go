package server_tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Activities Table Integration", func() {
	var userID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		// Insert into credentials first (to satisfy users FK)
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, "activitytest@example.com", "hashedpassword", time.Now(), time.Now())
		Expect(err).To(BeNil())

		// Insert into users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "activitytestuser", "activitytest@example.com", 0)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM activities")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve an activity for a user", func() {
		activityID := uuid.New()
		now := time.Now()
		message := "Test activity message"

		_, err := testDbInstance.Exec(`
			INSERT INTO activities (id, user_id, time, message)
			VALUES ($1, $2, $3, $4)
		`, activityID, userID, now, message)
		Expect(err).To(BeNil())

		// Query back
		row := testDbInstance.QueryRow(`
			SELECT id, user_id, time, message FROM activities WHERE user_id = $1
		`, userID)
		var gotID uuid.UUID
		var gotUserID uuid.UUID
		var gotTime time.Time
		var gotMsg string
		err = row.Scan(&gotID, &gotUserID, &gotTime, &gotMsg)
		Expect(err).To(BeNil())
		Expect(gotID).To(Equal(activityID))
		Expect(gotUserID).To(Equal(userID))
		Expect(gotMsg).To(Equal(message))
	})

	It("should enforce foreign key constraint on user_id", func() {
		nonExistentUserID := uuid.New()
		_, err := testDbInstance.Exec(`
			INSERT INTO activities (id, user_id, time, message)
			VALUES ($1, $2, $3, $4)
		`, uuid.New(), nonExistentUserID, time.Now(), "Should fail")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates foreign key constraint"))
	})
})
