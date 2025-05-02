package database_tests

import (
	"fmt"
	"rocket-backend/internal/database"
	"rocket-backend/internal/types"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Challenges Table", func() {
	var (
		srv database.Service
	)

	BeforeEach(func() {
		srv = database.NewWithConfig(connectionString)
	})

	AfterEach(func() {
		_, err := srv.ExecuteRawSQL("DELETE FROM user_challenges")
		Expect(err).NotTo(HaveOccurred())

		_, err = srv.ExecuteRawSQL("DELETE FROM challenges")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("InsertChallenge", func() {
		It("should insert a challenge successfully", func() {
			challenge := types.Challenge{
				ID:     uuid.New().String(),
				Text:   "Complete 10 push-ups",
				Points: 50,
			}

			err := srv.InsertChallenge(challenge)
			Expect(err).NotTo(HaveOccurred())

			challenges, err := srv.GetAllChallenges()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(challenges)).To(Equal(1))
			Expect(challenges[0].ID).To(Equal(challenge.ID))
			Expect(challenges[0].Text).To(Equal(challenge.Text))
			Expect(challenges[0].Points).To(Equal(challenge.Points))
		})
	})

	Context("GetAllChallenges", func() {
		It("should retrieve all challenges", func() {
			challenge1 := types.Challenge{
				ID:     uuid.New().String(),
				Text:   "Run 5 kilometers",
				Points: 100,
			}
			challenge2 := types.Challenge{
				ID:     uuid.New().String(),
				Text:   "Drink 2 liters of water",
				Points: 30,
			}

			err := srv.InsertChallenge(challenge1)
			Expect(err).NotTo(HaveOccurred())
			err = srv.InsertChallenge(challenge2)
			Expect(err).NotTo(HaveOccurred())

			challenges, err := srv.GetAllChallenges()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(challenges)).To(Equal(2))
		})
	})

	Context("AssignChallengesToUser", func() {
		It("should assign challenges to a user successfully", func() {
			userID := uuid.New()
			challenge1 := types.Challenge{
				ID:     "46849703-a389-4ecc-96cd-01d3a0c5ad15",
				Text:   "Meditate for 10 minutes",
				Points: 20,
			}
			challenge2 := types.Challenge{
				ID:     "f6a0c411-d992-4f5c-bf2c-79e100fb04d4",
				Text:   "Read 20 pages of a book",
				Points: 40,
			}

			err := srv.InsertChallenge(challenge1)
			Expect(err).NotTo(HaveOccurred())
			err = srv.InsertChallenge(challenge2)
			Expect(err).NotTo(HaveOccurred())

			err = srv.AssignChallengesToUser(userID, []types.Challenge{challenge1, challenge2})
			Expect(err).NotTo(HaveOccurred())

			userChallenges, err := srv.GetUserDailyChallenges(userID)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(userChallenges)).To(Equal(2))
			Expect(userChallenges[0].ID).To(Equal(challenge1.ID))
			Expect(userChallenges[1].ID).To(Equal(challenge2.ID))
		})
	})

	Context("CompleteChallenge", func() {
		It("should mark a challenge as completed", func() {
			userID := uuid.New()
			challenge := types.Challenge{
				ID:     uuid.New().String(),
				Text:   "Walk 10,000 steps",
				Points: 70,
			}

			err := srv.InsertChallenge(challenge)
			Expect(err).NotTo(HaveOccurred())

			err = srv.AssignChallengesToUser(userID, []types.Challenge{challenge})
			Expect(err).NotTo(HaveOccurred())

			dto := types.CompleteChallengesDTO{
				ChallengeID: uuid.MustParse(challenge.ID),
			}
			err = srv.CompleteChallenge(userID, dto)
			Expect(err).NotTo(HaveOccurred())

			userChallenges, err := srv.GetUserDailyChallenges(userID)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(userChallenges)).To(Equal(0))
		})
	})

	Context("ResetDailyChallenges", func() {
		It("should reset daily challenges for all users", func() {
			userID1 := uuid.New()
			userID2 := uuid.New()

			for i := range 5 {
				challenge := types.Challenge{
					ID:     uuid.New().String(),
					Text:   fmt.Sprintf("Challenge %d", i+1),
					Points: 10 * (i + 1),
				}
				err := srv.InsertChallenge(challenge)
				Expect(err).NotTo(HaveOccurred())
			}

			err := srv.AssignChallengesToUser(userID1, []types.Challenge{})
			Expect(err).NotTo(HaveOccurred())
			err = srv.AssignChallengesToUser(userID2, []types.Challenge{})
			Expect(err).NotTo(HaveOccurred())

			// Simulate that the challenges are from the previous day
			_, err = srv.ExecuteRawSQL(`UPDATE user_challenges SET date = date - INTERVAL '1 day';`)
			Expect(err).NotTo(HaveOccurred())

			err = srv.ResetDailyChallenges()
			Expect(err).NotTo(HaveOccurred())

			userChallenges1, err := srv.GetUserDailyChallenges(userID1)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(userChallenges1)).To(Equal(0))

			userChallenges2, err := srv.GetUserDailyChallenges(userID2)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(userChallenges2)).To(Equal(0))
		})
	})
})
