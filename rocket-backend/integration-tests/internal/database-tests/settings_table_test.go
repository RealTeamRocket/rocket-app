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
		_, err = srv.ExecuteRawSQL("DELETE FROM image_store")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Create Settings and retrieve them", func() {
		It("Create Settings and get settings", func() {
			id := uuid.New()
			userID := uuid.New()
			imageId := uuid.Nil
			stepGoal := 1

			settings := types.Settings{
				ID:       id,
				UserId:   userID,
				ImageId:  imageId,
				StepGoal: stepGoal,
			}

			err := srv.CreateSettings(settings)
			Expect(err).NotTo(HaveOccurred())

			savedSettings, err := srv.GetSettingsByUserID(userID)
			Expect(err).NotTo(HaveOccurred())

			Expect(settings.ID).To(Equal(savedSettings.ID))
			Expect(settings.UserId).To(Equal(savedSettings.UserId))
			Expect(settings.ImageId).To(Equal(savedSettings.ImageId))
			Expect(settings.StepGoal).To(Equal(savedSettings.StepGoal))
		})
	})

	Context("Update Settings", func() {
		It("Create Settings and get settings", func() {
			id := uuid.New()
			userID := uuid.New()
			imageId := uuid.Nil
			stepGoal := 1

			settings := types.Settings{
				ID:       id,
				UserId:   userID,
				ImageId:  imageId,
				StepGoal: stepGoal,
			}

			err := srv.CreateSettings(settings)
			Expect(err).NotTo(HaveOccurred())

			updatedSettings := types.SettingsDTO{
				StepGoal: 100,
			}
			imageId = uuid.New()

			err = srv.UpdateSettingsStepGoal(userID, updatedSettings.StepGoal)
			Expect(err).NotTo(HaveOccurred())

			err = srv.UpdateSettingsImage(userID, imageId)
			Expect(err).NotTo(HaveOccurred())

			savedSettings, err := srv.GetSettingsByUserID(userID)
			Expect(err).NotTo(HaveOccurred())

			Expect(settings.ID).To(Equal(savedSettings.ID))
			Expect(settings.UserId).To(Equal(savedSettings.UserId))
			Expect(updatedSettings.StepGoal).To(Equal(savedSettings.StepGoal))
			Expect(imageId).To(Equal(savedSettings.ImageId))
		})
	})

	Context("Update Step Goal", func() {
		It("should update the step goal independently", func() {
			userID := uuid.New()
			settings := types.Settings{
				ID:       uuid.New(),
				UserId:   userID,
				ImageId:  uuid.Nil,
				StepGoal: 5000,
			}

			err := srv.CreateSettings(settings)
			Expect(err).NotTo(HaveOccurred())

			err = srv.UpdateSettingsStepGoal(userID, 10000)
			Expect(err).NotTo(HaveOccurred())

			savedSettings, err := srv.GetSettingsByUserID(userID)
			Expect(err).NotTo(HaveOccurred())
			Expect(savedSettings.StepGoal).To(Equal(10000))
		})
	})

	Context("Update Image", func() {
		It("should update the image independently", func() {
			userID := uuid.New()
			settings := types.Settings{
				ID:       uuid.New(),
				UserId:   userID,
				ImageId:  uuid.Nil,
				StepGoal: 5000,
			}

			err := srv.CreateSettings(settings)
			Expect(err).NotTo(HaveOccurred())

			newImageID := uuid.New()
			err = srv.UpdateSettingsImage(userID, newImageID)
			Expect(err).NotTo(HaveOccurred())

			savedSettings, err := srv.GetSettingsByUserID(userID)
			Expect(err).NotTo(HaveOccurred())
			Expect(savedSettings.ImageId).To(Equal(newImageID))
		})
	})
})
