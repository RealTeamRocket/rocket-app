package server_tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Friends Table Integration", func() {
	var userID uuid.UUID
	var friendID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		friendID = uuid.New()
		now := time.Now()

		// Insert into credentials for both users
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10)
		`, userID, "user1@example.com", "hashedpassword", now, now,
			friendID, "user2@example.com", "hashedpassword", now, now)
		Expect(err).To(BeNil())

		// Insert into users for both users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4), ($5, $6, $7, $8)
		`, userID, "user1", "user1@example.com", 0,
			friendID, "user2", "user2@example.com", 10)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM friends")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should add and retrieve a friend", func() {
		_, err := testDbInstance.Exec(`
			INSERT INTO friends (user_id, friend_id)
			VALUES ($1, $2)
		`, userID, friendID)
		Expect(err).To(BeNil())

		rows, err := testDbInstance.Query(`
			SELECT friend_id FROM friends WHERE user_id = $1
		`, userID)
		Expect(err).To(BeNil())
		defer rows.Close()

		var found bool
		for rows.Next() {
			var fid uuid.UUID
			err := rows.Scan(&fid)
			Expect(err).To(BeNil())
			if fid == friendID {
				found = true
			}
		}
		Expect(found).To(BeTrue())
	})

	It("should not allow a user to friend themselves", func() {
		_, err := testDbInstance.Exec(`
			INSERT INTO friends (user_id, friend_id)
			VALUES ($1, $2)
		`, userID, userID)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates check constraint"))
	})

	It("should delete a friend", func() {
		_, err := testDbInstance.Exec(`
			INSERT INTO friends (user_id, friend_id)
			VALUES ($1, $2)
		`, userID, friendID)
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			DELETE FROM friends WHERE user_id = $1 AND friend_id = $2
		`, userID, friendID)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT COUNT(*) FROM friends WHERE user_id = $1 AND friend_id = $2
		`, userID, friendID)
		var count int
		err = row.Scan(&count)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(0))
	})

	It("should enforce foreign key constraint on friend_id", func() {
		nonExistentID := uuid.New()
		_, err := testDbInstance.Exec(`
			INSERT INTO friends (user_id, friend_id)
			VALUES ($1, $2)
		`, userID, nonExistentID)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates foreign key constraint"))
	})

	It("should retrieve followers correctly", func() {
		_, err := testDbInstance.Exec(`
			INSERT INTO friends (user_id, friend_id)
			VALUES ($1, $2)
		`, userID, friendID)
		Expect(err).To(BeNil())

		// Query followers of friendID (should include userID)
		rows, err := testDbInstance.Query(`
			SELECT user_id FROM friends WHERE friend_id = $1
		`, friendID)
		Expect(err).To(BeNil())
		defer rows.Close()

		var found bool
		for rows.Next() {
			var follower uuid.UUID
			err := rows.Scan(&follower)
			Expect(err).To(BeNil())
			if follower == userID {
				found = true
			}
		}
		Expect(found).To(BeTrue())

		// Query followers of userID (should be empty)
		rows, err = testDbInstance.Query(`
			SELECT user_id FROM friends WHERE friend_id = $1
		`, userID)
		Expect(err).To(BeNil())
		defer rows.Close()

		count := 0
		for rows.Next() {
			count++
		}
		Expect(count).To(Equal(0))
	})
})
