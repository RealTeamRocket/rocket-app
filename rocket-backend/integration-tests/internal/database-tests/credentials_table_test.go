package server_tests

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/google/uuid"
)

var _ = Describe("Credentials Table Integration", func() {
	var credsID uuid.UUID
	var email string

	BeforeEach(func() {
		credsID = uuid.New()
		email = "creduser@example.com"
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, credsID, email, "hashedpassword", time.Now(), time.Now())
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDbInstance.Exec("DELETE FROM credentials")
	})

	It("should insert and retrieve credentials by email", func() {
		row := testDbInstance.QueryRow(`
			SELECT id, email, password FROM credentials WHERE email = $1
		`, email)
		var gotID uuid.UUID
		var gotEmail, gotPassword string
		err := row.Scan(&gotID, &gotEmail, &gotPassword)
		Expect(err).To(BeNil())
		Expect(gotID).To(Equal(credsID))
		Expect(gotEmail).To(Equal(email))
		Expect(gotPassword).To(Equal("hashedpassword"))
	})

	It("should not allow duplicate emails", func() {
		_, err := testDbInstance.Exec(`
			INSERT INTO credentials (id, email, password, created_at, last_login)
			VALUES ($1, $2, $3, $4, $5)
		`, uuid.New(), email, "anotherpassword", time.Now(), time.Now())
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("duplicate key value"))
	})
})
