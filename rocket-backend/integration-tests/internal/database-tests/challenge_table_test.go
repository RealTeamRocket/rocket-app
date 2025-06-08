package server_tests

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Challenges Table Integration", func() {
	var userID uuid.UUID
	var challengeIDs []uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		// Insert into credentials first (to satisfy users FK)
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, "challengeuser@example.com", "hashedpassword", time.Now(), time.Now())
		Expect(err).To(BeNil())

		// Insert into users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "challengeuser", "challengeuser@example.com", 0)
		Expect(err).To(BeNil())

		// Insert some challenges
		challengeIDs = []uuid.UUID{uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New()}
		for i, cid := range challengeIDs {
			desc := fmt.Sprintf("Challenge %c", 'A'+i)
			_, err := testDbInstance.Exec(`
				INSERT INTO challenges (id, description, points_reward)
				VALUES ($1, $2, $3)
			`, cid, desc, 10*(i+1))
			Expect(err).To(BeNil())
		}
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM user_challenges")
		testDbInstance.Exec("DELETE FROM challenges")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should assign challenges to a user and retrieve them", func() {
		// Assign all challenges to user for today
		today := time.Now().Format("2006-01-02")
		for _, cid := range challengeIDs {
			_, err := testDbInstance.Exec(`
				INSERT INTO user_challenges (user_id, challenge_id, date, is_completed)
				VALUES ($1, $2, $3, FALSE)
			`, userID, cid, today)
			Expect(err).To(BeNil())
		}

		// Retrieve assigned challenges for user
		rows, err := testDbInstance.Query(`
			SELECT c.id, c.description, c.points_reward
			FROM user_challenges uc
			JOIN challenges c ON uc.challenge_id = c.id
			WHERE uc.user_id = $1 AND uc.date = $2 AND uc.is_completed = FALSE
		`, userID, today)
		Expect(err).To(BeNil())
		defer rows.Close()

		var count int
		for rows.Next() {
			var id uuid.UUID
			var desc string
			var points int
			err := rows.Scan(&id, &desc, &points)
			Expect(err).To(BeNil())
			count++
		}
		Expect(count).To(Equal(5))
	})

	It("should enforce foreign key constraint on challenge_id", func() {
		nonExistentChallengeID := uuid.New()
		today := time.Now().Format("2006-01-02")
		_, err := testDbInstance.Exec(`
			INSERT INTO user_challenges (user_id, challenge_id, date, is_completed)
			VALUES ($1, $2, $3, FALSE)
		`, userID, nonExistentChallengeID, today)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates foreign key constraint"))
	})
})
