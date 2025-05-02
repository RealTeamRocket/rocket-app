package database_tests

import (
"fmt"
"rocket-backend/internal/database"

"github.com/google/uuid"
. "github.com/onsi/ginkgo/v2"
. "github.com/onsi/gomega"
)

var _ = Describe("Friends table tests", func() {
var (
	srv database.Service
)

BeforeEach(func() {
	srv = database.NewWithConfig(connectionString)

	// Clean up tables before each test
	_, err := srv.ExecuteRawSQL("DELETE FROM friends")
	Expect(err).NotTo(HaveOccurred())
	_, err = srv.ExecuteRawSQL("DELETE FROM users")
	Expect(err).NotTo(HaveOccurred())
})

AfterEach(func() {
	// Clean up tables after each test
	_, err := srv.ExecuteRawSQL("DELETE FROM friends")
	Expect(err).NotTo(HaveOccurred())
	_, err = srv.ExecuteRawSQL("DELETE FROM users")
	Expect(err).NotTo(HaveOccurred())
})

Context("AddFriend", func() {
	It("should add a friend successfully", func() {
		userID := uuid.New()
		friendID := uuid.New()

		// Insert test users
		_, err := srv.ExecuteRawSQL(fmt.Sprintf(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ('%s', 'user1', 'user1@example.com', 100),
				   ('%s', 'friend1', 'friend1@example.com', 200)
		`, userID, friendID))
		Expect(err).NotTo(HaveOccurred())

		// Add friend
		err = srv.AddFriend(userID, friendID)
		Expect(err).NotTo(HaveOccurred())

		// Verify friend was added
		var count int
		row := srv.QueryRow(fmt.Sprintf(`
			SELECT COUNT(*) FROM friends WHERE user_id = '%s' AND friend_id = '%s'
		`, userID, friendID))
		err = row.Scan(&count)
		Expect(err).NotTo(HaveOccurred())
		Expect(count).To(Equal(1))
	})

	It("should not allow a user to add themselves as a friend", func() {
		userID := uuid.New()

		// Insert test user
		_, err := srv.ExecuteRawSQL(fmt.Sprintf(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ('%s', 'user1', 'user1@example.com', 100)
		`, userID))
		Expect(err).NotTo(HaveOccurred())

		// Attempt to add self as friend
		err = srv.AddFriend(userID, userID)
		Expect(err).To(HaveOccurred())
	})
})

Context("GetFriends", func() {
	It("should retrieve a list of friends for a user", func() {
		userID := uuid.New()
		friendID1 := uuid.New()
		friendID2 := uuid.New()

		// Insert test users
		_, err := srv.ExecuteRawSQL(fmt.Sprintf(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ('%s', 'user1', 'user1@example.com', 100),
				   ('%s', 'friend1', 'friend1@example.com', 200),
				   ('%s', 'friend2', 'friend2@example.com', 300)
		`, userID, friendID1, friendID2))
		Expect(err).NotTo(HaveOccurred())

		// Add friends
		err = srv.AddFriend(userID, friendID1)
		Expect(err).NotTo(HaveOccurred())
		err = srv.AddFriend(userID, friendID2)
		Expect(err).NotTo(HaveOccurred())

		// Retrieve friends
		friends, err := srv.GetFriends(userID)
		Expect(err).NotTo(HaveOccurred())
		Expect(friends).To(HaveLen(2))

		// Verify friends are sorted by username
		Expect(friends[0].Username).To(Equal("friend1"))
		Expect(friends[1].Username).To(Equal("friend2"))
	})

	It("should return an error if the user has no friends", func() {
		userID := uuid.New()

		// Insert test user
		_, err := srv.ExecuteRawSQL(fmt.Sprintf(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ('%s', 'user1', 'user1@example.com', 100)
		`, userID))
		Expect(err).NotTo(HaveOccurred())

		// Attempt to retrieve friends
		_, err = srv.GetFriends(userID)
		Expect(err).To(HaveOccurred())
	})
})

Context("GetFriendsRankedByPoints", func() {
	It("should return an error if the user has no friends", func() {
		userID := uuid.New()

		// Insert test user
		_, err := srv.ExecuteRawSQL(fmt.Sprintf(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ('%s', 'user1', 'user1@example.com', 100)
		`, userID))
		Expect(err).NotTo(HaveOccurred())

		// Attempt to retrieve friends ranked by points
		_, err = srv.GetFriendsRankedByPoints(userID)
		Expect(err).To(HaveOccurred())
	})
})
})
