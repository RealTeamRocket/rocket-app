package server_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"

	"rocket-backend/integration-tests/mocks"
	"rocket-backend/internal/server"
	"rocket-backend/internal/types"
)

var _ = Describe("PublicHandlers", func() {
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

		router.GET("/hello", srv.HelloWorldHandler)
		router.GET("/health", srv.HealthHandler)
		router.POST("/login", srv.LoginHandler)
		router.POST("/register", srv.RegisterHandler)
	})

	Describe("HelloWorldHandler", func() {
		It("should return a hello world message", func() {
			req, err := http.NewRequest("GET", "/hello", nil)
			Expect(err).To(BeNil())

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(recorder.Body.String()).To(ContainSubstring("Hello World"))
		})
	})

	Describe("HealthHandler", func() {
		It("should return health status", func() {
			mock.HealthFunc = func() map[string]string {
				return map[string]string{"status": "up"}
			}

			req, err := http.NewRequest("GET", "/health", nil)
			Expect(err).To(BeNil())

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(recorder.Body.String()).To(ContainSubstring(`"status":"up"`))
		})
	})

	Describe("LoginHandler", func() {
		Context("with valid credentials", func() {
			BeforeEach(func() {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
				mock.GetUserByEmailFunc = func(email string) (types.Credentials, error) {
					return types.Credentials{
						ID:       uuid.New(),
						Email:    email,
						Password: string(hashedPassword),
					}, nil
				}
			})

			It("should return a token", func() {
				body := `{"email": "test@example.com", "password": "password"}`
				req, err := http.NewRequest("POST", "/login", strings.NewReader(body))
				Expect(err).To(BeNil())
				req.Header.Set("Content-Type", "application/json")

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))
				Expect(recorder.Body.String()).To(ContainSubstring("token"))
			})
		})

		Context("with invalid credentials", func() {
			BeforeEach(func() {
				mock.GetUserByEmailFunc = func(email string) (types.Credentials, error) {
					return types.Credentials{}, errors.New("user not found")
				}
			})

			It("should return an error", func() {
				body := `{"email": "test@example.com", "password": "wrongpassword"}`
				req, err := http.NewRequest("POST", "/login", strings.NewReader(body))
				Expect(err).To(BeNil())
				req.Header.Set("Content-Type", "application/json")

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
				Expect(recorder.Body.String()).To(ContainSubstring("Invalid username or password"))
			})
		})
	})

	Describe("RegisterHandler", func() {
		Context("successful registration", func() {
			BeforeEach(func() {
				mock.CheckEmailFunc = func(email string) error {
					return nil
				}
				mock.SaveCredentialsFunc = func(creds types.Credentials) error {
					return nil
				}
				mock.SaveUserProfileFunc = func(user types.User) error {
					return nil
				}
			})

			It("should register a new user", func() {
				body := `{"email": "test@example.com", "password": "password", "username": "testuser"}`
				req, err := http.NewRequest("POST", "/register", strings.NewReader(body))
				Expect(err).To(BeNil())
				req.Header.Set("Content-Type", "application/json")

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusOK))
				Expect(recorder.Body.String()).To(ContainSubstring("User registered successfully"))
			})
		})

		Context("email already exists", func() {
			BeforeEach(func() {
				mock.CheckEmailFunc = func(email string) error {
					return errors.New("email already exists")
				}
			})

			It("should return an error", func() {
				body := `{"email": "test@example.com", "password": "password", "username": "testuser"}`
				req, err := http.NewRequest("POST", "/register", strings.NewReader(body))
				Expect(err).To(BeNil())
				req.Header.Set("Content-Type", "application/json")

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, req)

				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
				Expect(recorder.Body.String()).To(ContainSubstring("Email already exists"))
			})
		})
	})
})
