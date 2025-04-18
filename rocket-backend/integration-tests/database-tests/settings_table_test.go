package database_tests

import (
	"rocket-backend/internal/database"
	// "rocket-backend/internal/types"

	// "github.com/google/uuid"
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
	})

	Context("Update Settings", func() {
	})
})
