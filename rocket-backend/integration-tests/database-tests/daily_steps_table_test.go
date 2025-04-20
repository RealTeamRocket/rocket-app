package database_tests

import (
	"rocket-backend/internal/database"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Daily Steps Table", func() {
	var (
		srv database.Service
	)

	BeforeEach(func() {
		srv = database.NewWithConfig(connectionString)
	})

	AfterEach(func() {
		_, err := srv.ExecuteRawSQL("DELETE FROM daily_steps")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("UpdateDailySteps", func() {
		It("should insert daily steps successfully", func() {
			userID := uuid.New()
			steps := 10000
			currentDate := time.Now().Format("2006-01-02")

			// Insert daily steps
			err := srv.UpdateDailySteps(userID, steps)
			Expect(err).NotTo(HaveOccurred())

			// Verify that the steps were inserted correctly
			var savedSteps int
			err = srv.QueryRow("SELECT steps_taken FROM daily_steps WHERE user_id = $1 AND date = $2", userID, currentDate).Scan(&savedSteps)
			Expect(err).NotTo(HaveOccurred())
			Expect(savedSteps).To(Equal(steps))
		})
	})
})
