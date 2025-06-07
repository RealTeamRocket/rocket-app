package database_tests

import (
	"rocket-backend/internal/database"
	"rocket-backend/internal/types"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Users Table", func() {
	var (
		srv database.Service
	)

	BeforeEach(func() {
		srv = database.NewWithConfig(connectionString)
	})

	AfterEach(func() {
		_, err := srv.ExecuteRawSQL("DELETE FROM users")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("SaveUserProfile", func() {
		It("should save user profile successfully", func() {
			id := uuid.New()
			username := "johndoe"
			email := "john@doe.com"
			rocketPoints := 100

			user := types.User{
				ID:           id,
				Username:     username,
				Email:        email,
				RocketPoints: rocketPoints,
			}

			err := srv.SaveUserProfile(user)
			Expect(err).NotTo(HaveOccurred())

			// Verify that the user profile was saved correctly
			var savedUser types.User
			savedUser, err = srv.GetUserByID(user.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(savedUser.ID).To(Equal(user.ID))
			Expect(savedUser.Username).To(Equal(user.Username))
			Expect(savedUser.Email).To(Equal(user.Email))
			Expect(savedUser.RocketPoints).To(Equal(user.RocketPoints))
		})
	})
})
