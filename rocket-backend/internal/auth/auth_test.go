package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth Service", func() {
	var (
		jwtSecret   string
		authService *AuthService
		userID      uuid.UUID
	)

	BeforeEach(func() {
		jwtSecret = "testsecret"
		authService = NewAuthService(jwtSecret)
		userID = uuid.New()
	})

	Describe("GenerateToken", func() {
		It("should generate a valid token", func() {
			tokenString, err := authService.GenerateToken(userID)
			Expect(err).NotTo(HaveOccurred())
			Expect(tokenString).NotTo(BeEmpty())
		})
	})

	Describe("ParseToken", func() {
		It("should parse a valid token", func() {
			tokenString, err := authService.GenerateToken(userID)
			Expect(err).NotTo(HaveOccurred())

			token, err := authService.ParseToken(tokenString)
			Expect(err).NotTo(HaveOccurred())
			Expect(token).NotTo(BeNil())
			Expect(token.Valid).To(BeTrue())
		})
	})

	Describe("ValidateToken", func() {
		It("should validate a token and extract the user ID", func() {
			tokenString, err := authService.GenerateToken(userID)
			Expect(err).NotTo(HaveOccurred())

			token, err := authService.ParseToken(tokenString)
			Expect(err).NotTo(HaveOccurred())
			Expect(token).NotTo(BeNil())
			Expect(token.Valid).To(BeTrue())

			parsedUserID, err := authService.ValidateToken(token)
			Expect(err).NotTo(HaveOccurred())
			Expect(parsedUserID).To(Equal(userID))
		})
	})

	Describe("InvalidToken", func() {
		It("should return an error for an invalid token", func() {
			invalidTokenString := "invalid.token.string"

			token, err := authService.ParseToken(invalidTokenString)
			Expect(err).To(HaveOccurred())
			Expect(token).NotTo(BeNil())
			Expect(token.Valid).To(BeFalse())
		})
	})

	Describe("ExpiredToken", func() {
		It("should return an error for an expired token", func() {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user_id": userID.String(),
				"exp":     time.Now().Add(time.Second * 1).Unix(),
			})

			tokenString, err := token.SignedString([]byte(jwtSecret))
			Expect(err).NotTo(HaveOccurred())

			// Wait for the token to expire
			time.Sleep(time.Second * 2)

			parsedToken, err := authService.ParseToken(tokenString)
			Expect(err).To(HaveOccurred())
			Expect(parsedToken).NotTo(BeNil())
			Expect(parsedToken.Valid).To(BeFalse())
		})
	})
})

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}
