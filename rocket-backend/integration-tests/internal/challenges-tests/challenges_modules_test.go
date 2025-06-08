package challenges_tests

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"rocket-backend/internal/challenges"
	"rocket-backend/internal/database"
	"github.com/google/uuid"
)

var _ = Describe("Challenges Module", func() {
	var (
		dbService database.Service
		manager   *challenges.ChallengeManager
		userID    uuid.UUID
	)

	BeforeEach(func() {
		dbService = database.NewWithConfig(connectionString)
		manager = challenges.NewChallengeManager(dbService)
		userID = uuid.New()

		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, NOW(), NOW())
		`, userID, "challengeuser@example.com", "hashedpassword")
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "challengeuser", "challengeuser@example.com", 0)
		Expect(err).To(BeNil())
	})

	It("should load challenges from file", func() {
		wd, _ := os.Getwd()
		filePath := filepath.Join(wd, "../../../internal/challenges/challenges.json")
		chals, err := challenges.LoadChallengesFromFile(filePath)
		Expect(err).To(BeNil())
		Expect(len(chals)).To(BeNumerically(">=", 1))
	})

	It("should assign daily challenges to a user", func() {
		dailies, err := manager.GetDailies(userID)
		Expect(err).To(BeNil())
		Expect(len(dailies)).To(Equal(5))
	})
})
