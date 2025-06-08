package server_tests

import (
	"bytes"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Image Store Table Integration", func() {
	var userID uuid.UUID
	var imageID uuid.UUID

	BeforeEach(func() {
		userID = uuid.New()
		now := time.Now()
		// Insert into credentials first (to satisfy users FK)
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, "imguser@example.com", "hashedpassword", now, now)
		Expect(err).To(BeNil())

		// Insert into users
		_, err = testDbInstance.Exec(`
			INSERT INTO users (id, username, email, rocketpoints)
			VALUES ($1, $2, $3, $4)
		`, userID, "imguser", "imguser@example.com", 0)
		Expect(err).To(BeNil())

		// Insert into settings (needed for GetUserImage)
		_, err = testDbInstance.Exec(`
			INSERT INTO settings (id, user_id, step_goal)
			VALUES ($1, $2, $3)
		`, uuid.New(), userID, 10000)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM settings")
		testDbInstance.Exec("DELETE FROM image_store")
		testDbInstance.Exec("DELETE FROM users")
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve an image", func() {
		imageID = uuid.New()
		imgName := "test.png"
		imgData := []byte{0x89, 0x50, 0x4E, 0x47} // PNG header

		_, err := testDbInstance.Exec(`
			INSERT INTO image_store (id, image_name, image_data)
			VALUES ($1, $2, $3)
		`, imageID, imgName, imgData)
		Expect(err).To(BeNil())

		// Update settings to reference this image
		_, err = testDbInstance.Exec(`
			UPDATE settings SET image_id = $1 WHERE user_id = $2
		`, imageID, userID)
		Expect(err).To(BeNil())

		// Query back
		row := testDbInstance.QueryRow(`
			SELECT i.id, i.image_name, i.image_data
			FROM settings s
			JOIN image_store i ON s.image_id = i.id
			WHERE s.user_id = $1
		`, userID)
		var gotID uuid.UUID
		var gotName string
		var gotData []byte
		err = row.Scan(&gotID, &gotName, &gotData)
		Expect(err).To(BeNil())
		Expect(gotID).To(Equal(imageID))
		Expect(gotName).To(Equal(imgName))
		Expect(bytes.Equal(gotData, imgData)).To(BeTrue())
	})

	It("should return no rows if user has no image", func() {
		row := testDbInstance.QueryRow(`
			SELECT i.id, i.image_name, i.image_data
			FROM settings s
			JOIN image_store i ON s.image_id = i.id
			WHERE s.user_id = $1
		`, userID)
		var gotID uuid.UUID
		var gotName string
		var gotData []byte
		err := row.Scan(&gotID, &gotName, &gotData)
		Expect(err).ToNot(BeNil())
	})
})
