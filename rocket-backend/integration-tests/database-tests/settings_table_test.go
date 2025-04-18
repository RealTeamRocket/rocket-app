package database_tests

import (
	"rocket-backend/internal/database"
	"rocket-backend/internal/types"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Settings table tests", func() {
	var (
		srv database.Service
	)

	BeforeEach(func() {
		srv = database.NewWithConfig(connectionString)
	})

	AfterEach(func() {
		_, err := srv.ExecuteRawSQL("DELETE FROM settings")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Create Settings and retrieve them", func() {
		It("should create new settings and get the out of db", func() {
			id := uuid.New()
			userID := uuid.New()
			stepGoal := 10000
			profileImage := "Dummy png"

			settings := types.Settings {
				ID:           id,
				UserId: userID,
				StepGoal: stepGoal,
				ProfileImage: profileImage,
			}

			err := srv.CreateSettings(settings)
			Expect(err).NotTo(HaveOccurred())

			savedSettings, err := srv.GetSettingsByUserID(userID)
			Expect(err).NotTo(HaveOccurred())

			Expect(settings.ID).To(Equal(savedSettings.ID))
			Expect(settings.UserId).To(Equal(savedSettings.UserId))
			Expect(settings.StepGoal).To(Equal(savedSettings.StepGoal))
			Expect(settings.ProfileImage).To(Equal(savedSettings.ProfileImage))
		})
	})

	Context("Update Settings", func() {
		It("should update created settings", func() {
			id := uuid.New()
			userID := uuid.New()
			stepGoal := 10000
			profileImage := "Dummy png"

			settings := types.Settings {
				ID:           id,
				UserId: userID,
				StepGoal: stepGoal,
				ProfileImage: profileImage,
			}

			err := srv.CreateSettings(settings)
			Expect(err).NotTo(HaveOccurred())

			updateSteps := 1
			updateProfileImage := "updated Image"

			updateSettings := types.SettingsDTO {
				StepGoal: updateSteps,
				ProfileImage: updateProfileImage,
			}

			err = srv.UpdateSettings(userID, updateSettings)
			Expect(err).NotTo(HaveOccurred())

			savedSettings, err := srv.GetSettingsByUserID(userID)
			Expect(err).NotTo(HaveOccurred())

			Expect(settings.ID).To(Equal(savedSettings.ID))
			Expect(settings.UserId).To(Equal(savedSettings.UserId))
			Expect(updateSettings.StepGoal).To(Equal(savedSettings.StepGoal))
			Expect(updateSettings.ProfileImage).To(Equal(savedSettings.ProfileImage))
		})
	})
})
