package server_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"rocket-backend/integration-tests/mocks"
	"rocket-backend/internal/auth"
	"rocket-backend/internal/server"
	"rocket-backend/internal/types"
)

var _ = Describe("AuthMiddleware", func() {
	var (
		srv    *server.Server
		router *gin.Engine
		mock   *mocks.MockDB
	)

	BeforeEach(func() {
		mock = &mocks.MockDB{}
		srv = server.NewServerWithDB(mock, 8080, "test-secret")

		gin.SetMode(gin.TestMode)
		router = gin.New()

		router.Use(srv.AuthMiddleware())
		router.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "You are authorized"})
		})
	})

	Context("when the Authorization header is missing", func() {
		It("should return an unauthorized error", func() {
			req, err := http.NewRequest("GET", "/protected", nil)
			Expect(err).To(BeNil())

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
			Expect(recorder.Body.String()).To(ContainSubstring("Authorization header required"))
		})
	})

	Context("when the token is invalid", func() {
		It("should return an unauthorized error", func() {
			req, err := http.NewRequest("GET", "/protected", nil)
			Expect(err).To(BeNil())
			req.Header.Set("Authorization", "Bearer invalidtoken")

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
			Expect(recorder.Body.String()).To(ContainSubstring("Invalid token"))
		})
	})

	Context("when the token is valid but the user does not exist", func() {
		BeforeEach(func() {
			mock.GetUserByIDFunc = func(userID uuid.UUID) (types.User, error) {
				return types.User{}, errors.New("user not found")
			}
		})

		It("should return an unauthorized error", func() {
			authService := auth.NewAuthService("test-secret")
			tokenString, _ := authService.GenerateToken(uuid.New())

			req, err := http.NewRequest("GET", "/protected", nil)
			Expect(err).To(BeNil())
			req.Header.Set("Authorization", "Bearer "+tokenString)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
			Expect(recorder.Body.String()).To(ContainSubstring("User does not exist"))
		})
	})

	Context("when the token is valid and the user exists", func() {
		BeforeEach(func() {
			mock.GetUserByIDFunc = func(userID uuid.UUID) (types.User, error) {
				return types.User{
					ID:       userID,
					Username: "testuser",
				}, nil
			}
		})

		It("should allow access to the protected route", func() {
			authService := auth.NewAuthService("test-secret")
			tokenString, _ := authService.GenerateToken(uuid.New())

			req, err := http.NewRequest("GET", "/protected", nil)
			Expect(err).To(BeNil())
			req.Header.Set("Authorization", "Bearer "+tokenString)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(recorder.Body.String()).To(ContainSubstring("You are authorized"))
		})
	})
})
