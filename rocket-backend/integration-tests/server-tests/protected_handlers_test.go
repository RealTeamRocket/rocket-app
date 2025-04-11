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
	"rocket-backend/internal/server"
	"rocket-backend/internal/types"
)

var _ = Describe("ProtectedHandler", func() {
    var (
        srv      *server.Server
        router   *gin.Engine
        mock     *mocks.MockDB
        userUUID uuid.UUID
    )

    BeforeEach(func() {
        mock = &mocks.MockDB{}
        userUUID = uuid.New()

        srv = server.NewServerWithDB(mock, 8080, "test-secret")

        gin.SetMode(gin.TestMode)
        router = gin.New()

        router.Use(func(c *gin.Context) {
            c.Set("userID", userUUID.String())
            c.Next()
        })

        router.GET("/auth-hello", srv.AuthHelloHandler)
        router.GET("/authenticated", srv.Authenticated)
        router.POST("/update-steps", srv.UpdateSteps)
    })

    Describe("AuthHelloHandler", func() {
        Context("when the user is authenticated and valid", func() {
            BeforeEach(func() {
                mock.GetUserByIDFunc = func(userID uuid.UUID) (types.User, error) {
                    return types.User{
                        Username: "testuser",
                    }, nil
                }
            })

            It("should return a greeting with the username", func() {
                req, err := http.NewRequest("GET", "/auth-hello", nil)
                Expect(err).To(BeNil())

                recorder := httptest.NewRecorder()
                router.ServeHTTP(recorder, req)

                Expect(recorder.Code).To(Equal(http.StatusOK))
                Expect(recorder.Body.String()).To(ContainSubstring("Hello from auth"))
                Expect(recorder.Body.String()).To(ContainSubstring("testuser"))
            })
        })

        Context("when userID is missing", func() {
            It("should return an unauthorized error", func() {
                noUserRouter := gin.New()
                noUserRouter.GET("/auth-hello", srv.AuthHelloHandler)

                req, err := http.NewRequest("GET", "/auth-hello", nil)
                Expect(err).To(BeNil())

                recorder := httptest.NewRecorder()
                noUserRouter.ServeHTTP(recorder, req)
                Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
                Expect(recorder.Body.String()).To(ContainSubstring("User not authenticated"))
            })
        })

        Context("when GetUserByID returns an error", func() {
            BeforeEach(func() {
                mock.GetUserByIDFunc = func(userID uuid.UUID) (types.User, error) {
                    return types.User{}, errors.New("db error")
                }
            })

            It("should return an internal server error", func() {
                req, err := http.NewRequest("GET", "/auth-hello", nil)
                Expect(err).To(BeNil())

                recorder := httptest.NewRecorder()
                router.ServeHTTP(recorder, req)
                Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
                Expect(recorder.Body.String()).To(ContainSubstring("Failed to retrieve user"))
            })
        })
    })
})
