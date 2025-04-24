package database_tests

import (
	"fmt"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"rocket-backend/internal/database"
)

var _ = Describe("Image store table tests", func() {
	var (
		srv database.Service
	)

	BeforeEach(func() {
		srv = database.NewWithConfig(connectionString)
	})

	AfterEach(func() {
		// Clean up image_store table
		_, err := srv.ExecuteRawSQL("DELETE FROM image_store")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("SaveImage and verify storage", func() {
		It("should save image and be retrievable via raw SQL count", func() {
			filename := "test_image.png"
			data := []byte{0x01, 0x02, 0x03, 0x04}

			// Save the image
			id, err := srv.SaveImage(filename, data)
			Expect(err).NotTo(HaveOccurred())
			Expect(id).NotTo(Equal(uuid.Nil))

			// Verify the row exists via raw SQL
			rawQuery := fmt.Sprintf("SELECT * FROM image_store WHERE id = '%s'", id)
			res, err := srv.ExecuteRawSQL(rawQuery)
			Expect(err).NotTo(HaveOccurred())

			// Exec on SELECT returns RowsAffected; expect 1 row
			rows, err := res.RowsAffected()
			Expect(err).NotTo(HaveOccurred())
			Expect(rows).To(Equal(int64(1)))
		})
	})
})
