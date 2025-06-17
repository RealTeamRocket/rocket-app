package server_tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Users Table Integration", func() {
	var userID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		now := time.Now()
		// Insert into credentials first (to satisfy users FK)
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, "useruser@example.com", "hashedpassword", now, now)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve a user", func() {
		username := "useruser"
		email := "useruser@example.com"
		rocketpoints := 42

		_, err := testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, username, email, rocketpoints)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT id, username, email, rocketpoints FROM users WHERE id = $1
		`, userID)
		var gotID uuid.UUID
		var gotUsername, gotEmail string
		var gotPoints int
		err = row.Scan(&gotID, &gotUsername, &gotEmail, &gotPoints)
		Expect(err).To(BeNil())
		Expect(gotID).To(Equal(userID))
		Expect(gotUsername).To(Equal(username))
		Expect(gotEmail).To(Equal(email))
		Expect(gotPoints).To(Equal(rocketpoints))
	})

	It("should update rocketpoints for a user", func() {
		_, err := testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "useruser", "useruser@example.com", 0)
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			UPDATE users SET rocketpoints = rocketpoints + $1 WHERE id = $2
		`, 10, userID)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT rocketpoints FROM users WHERE id = $1
		`, userID)
		var gotPoints int
		err = row.Scan(&gotPoints)
		Expect(err).To(BeNil())
		Expect(gotPoints).To(Equal(10))
	})

	It("should enforce foreign key constraint on id", func() {
		nonExistentCredID := uuid.New()
		_, err := testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, nonExistentCredID, "nouser", "nouser@example.com", 0)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates foreign key constraint"))
	})

	It("should update and retrieve user name", func() {
		username := "useruser"
		email := "useruser@example.com"
		rocketpoints := 42

		_, err := testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, username, email, rocketpoints)
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			UPDATE users SET username = $1 WHERE id = $2
		`, "newname", userID)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT username FROM users WHERE id = $1
		`, userID)
		var gotUsername string
		err = row.Scan(&gotUsername)
		Expect(err).To(BeNil())
		Expect(gotUsername).To(Equal("newname"))
	})

	It("should update and retrieve user email in both users and credentials", func() {
		username := "useruser"
		email := "useruser@example.com"
		rocketpoints := 42

		_, err := testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, username, email, rocketpoints)
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			UPDATE users SET email = $1 WHERE id = $2
		`, "newemail@example.com", userID)
		Expect(err).To(BeNil())
		_, err = testDbInstance.Exec(`
			UPDATE credentials SET email = $1 WHERE id = $2
		`, "newemail@example.com", userID)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT email FROM users WHERE id = $1
		`, userID)
		var gotEmail string
		err = row.Scan(&gotEmail)
		Expect(err).To(BeNil())
		Expect(gotEmail).To(Equal("newemail@example.com"))

		row = testDbInstance.QueryRow(`
			SELECT email FROM credentials WHERE id = $1
		`, userID)
		err = row.Scan(&gotEmail)
		Expect(err).To(BeNil())
		Expect(gotEmail).To(Equal("newemail@example.com"))
	})

	It("should get user ID by username", func() {
		username := "useruser"
		email := "useruser@example.com"
		rocketpoints := 42

		_, err := testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, username, email, rocketpoints)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT id FROM users WHERE username = $1
		`, username)
		var gotID uuid.UUID
		err = row.Scan(&gotID)
		Expect(err).To(BeNil())
		Expect(gotID).To(Equal(userID))
	})

	It("should delete a user and its credentials", func() {
		username := "useruser"
		email := "useruser@example.com"
		rocketpoints := 42

		_, err := testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, username, email, rocketpoints)
		Expect(err).To(BeNil())

		_, err = testDbInstance.Exec(`
			DELETE FROM users WHERE id = $1
		`, userID)
		Expect(err).To(BeNil())
		_, err = testDbInstance.Exec(`
			DELETE FROM credentials WHERE id = $1
		`, userID)
		Expect(err).To(BeNil())

		row := testDbInstance.QueryRow(`
			SELECT COUNT(*) FROM users WHERE id = $1
		`, userID)
		var count int
		err = row.Scan(&count)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(0))
	})
})
