package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Public Handlers", func() {
	var (
		s  *Server
		r  *gin.Engine
		rr *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		s = &Server{}
		r = gin.New()
		api := r.Group("/api/v1")
		{
			api.GET("/", s.HelloWorldHandler)
		}
		rr = httptest.NewRecorder()
	})

	Describe("HelloWorldHandler", func() {
		It("should return Hello World message", func() {
			// Create a test HTTP request
			req, err := http.NewRequest("GET", "/api/v1/", nil)
			Expect(err).NotTo(HaveOccurred())

			// Serve the HTTP request
			r.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))

			// Check the response body
			expected := `{"message":"Hello World"}`
			Expect(rr.Body.String()).To(MatchJSON(expected))
		})
	})
})

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Suite")
}
