package server_tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Chat Messages Table Integration", func() {
	var userID uuid.UUID
	var otherUserID uuid.UUID
	var messageID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		otherUserID = uuid.New()
		now := time.Now()

		// Insert into credentials for both users
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10)
		`, userID, "chatuser@example.com", "hashedpassword", now, now,
			otherUserID, "otheruser@example.com", "hashedpassword", now, now)
		Expect(err).To(BeNil())

		// Insert into users for both users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4), ($5, $6, $7, $8)
		`, userID, "chatuser", "chatuser@example.com", 0,
			otherUserID, "otheruser", "otheruser@example.com", 0)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM chat_messages_reactions")
		testDbInstance.Exec("DELETE FROM chat_messages")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve a chat message", func() {
		now := time.Now()
		messageID = uuid.New()
		message := "Hello, world!"

		_, err := testDbInstance.Exec(`
			INSERT INTO chat_messages (id, user_id, message, timestamp)
			VALUES ($1, $2, $3, $4)
		`, messageID, userID, message, now)
		Expect(err).To(BeNil())

		// Query back
		row := testDbInstance.QueryRow(`
			SELECT id, user_id, message, timestamp FROM chat_messages WHERE id = $1
		`, messageID)
		var gotID uuid.UUID
		var gotUserID uuid.UUID
		var gotMsg string
		var gotTime time.Time
		err = row.Scan(&gotID, &gotUserID, &gotMsg, &gotTime)
		Expect(err).To(BeNil())
		Expect(gotID).To(Equal(messageID))
		Expect(gotUserID).To(Equal(userID))
		Expect(gotMsg).To(Equal(message))
	})

	It("should enforce foreign key constraint on user_id", func() {
		nonExistentUserID := uuid.New()
		_, err := testDbInstance.Exec(`
			INSERT INTO chat_messages (id, user_id, message, timestamp)
			VALUES ($1, $2, $3, $4)
		`, uuid.New(), nonExistentUserID, "Should fail", time.Now())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("violates foreign key constraint"))
	})

	It("should allow reactions and count them", func() {
		now := time.Now()
		messageID = uuid.New()
		// Insert a message from userID
		_, err := testDbInstance.Exec(`
			INSERT INTO chat_messages (id, user_id, message, timestamp)
			VALUES ($1, $2, $3, $4)
		`, messageID, userID, "React to me!", now)
		Expect(err).To(BeNil())

		// Insert a reaction from otherUserID
		_, err = testDbInstance.Exec(`
			INSERT INTO chat_messages_reactions (message_id, user_id, created_at)
			VALUES ($1, $2, $3)
		`, messageID, otherUserID, now)
		Expect(err).To(BeNil())

		// Count reactions
		row := testDbInstance.QueryRow(`
			SELECT COUNT(*) FROM chat_messages_reactions WHERE message_id = $1
		`, messageID)
		var count int
		err = row.Scan(&count)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(1))
	})

	It("should enforce unique reaction per user per message", func() {
		now := time.Now()
		messageID = uuid.New()
		// Insert a message from userID
		_, err := testDbInstance.Exec(`
			INSERT INTO chat_messages (id, user_id, message, timestamp)
			VALUES ($1, $2, $3, $4)
		`, messageID, userID, "React to me!", now)
		Expect(err).To(BeNil())

		// Insert a reaction from otherUserID
		_, err = testDbInstance.Exec(`
			INSERT INTO chat_messages_reactions (message_id, user_id, created_at)
			VALUES ($1, $2, $3)
		`, messageID, otherUserID, now)
		Expect(err).To(BeNil())

		// Try to insert a duplicate reaction
		_, err = testDbInstance.Exec(`
			INSERT INTO chat_messages_reactions (message_id, user_id, created_at)
			VALUES ($1, $2, $3)
		`, messageID, otherUserID, now)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("unique_message_user"))
	})
})
