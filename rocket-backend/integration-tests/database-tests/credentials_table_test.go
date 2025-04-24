package database_tests

import (
	"rocket-backend/internal/database"
	"rocket-backend/internal/types"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Credentials Table", func() {
	var (
		srv database.Service
	)

	BeforeEach(func() {
		srv = database.NewWithConfig(connectionString)
	})

	AfterEach(func() {
		_, err := srv.ExecuteRawSQL("DELETE FROM credentials")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("SaveCredentials", func() {
		It("should save credentials successfully", func() {
			id := uuid.New()
			email := "john@doe.com"
			password := "securepassword"
			createdAt := time.Now().Format(time.RFC3339)
			lastLogin := time.Now().Format(time.RFC3339)

			credentials := types.Credentials{
				ID:        id,
				Email:     email,
				Password:  password,
				CreatedAt: createdAt,
				LastLogin: lastLogin,
			}

			err := srv.SaveCredentials(credentials)
			Expect(err).NotTo(HaveOccurred())

			// Verify that the credentials were saved correctly
			savedCreds, err := srv.GetUserByEmail(email)
			Expect(err).NotTo(HaveOccurred())
			Expect(savedCreds.ID).To(Equal(credentials.ID))
			Expect(savedCreds.Email).To(Equal(credentials.Email))
			Expect(savedCreds.Password).To(Equal(credentials.Password))
			Expect(savedCreds.CreatedAt).To(Equal(credentials.CreatedAt))
			Expect(savedCreds.LastLogin).To(Equal(credentials.LastLogin))
		})
	})

	Context("CheckEmail", func() {
		It("should return an error if email already exists", func() {
			id := uuid.New()
			email := "john@doe.com"
			password := "securepassword"
			createdAt := time.Now().Format(time.RFC3339)
			lastLogin := time.Now().Format(time.RFC3339)

			credentials := types.Credentials{
				ID:        id,
				Email:     email,
				Password:  password,
				CreatedAt: createdAt,
				LastLogin: lastLogin,
			}

			err := srv.SaveCredentials(credentials)
			Expect(err).NotTo(HaveOccurred())

			err = srv.CheckEmail(email)
			Expect(err).To(HaveOccurred())
		})
	})
})
