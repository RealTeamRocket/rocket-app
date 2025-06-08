package server_tests

import (
	"net/http"
	// "os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthMiddleware", func() {
	var token string

	BeforeEach(func() {
		token = registerAndLogin("middleware@example.com", "password123", "middlewareuser")
	})

	It("should reject requests without Authorization header or cookie", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/user", nil)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(401))
	})

	It("should reject requests with invalid token", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/user", nil)
		req.Header.Set("Authorization", "Bearer invalidtoken")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(401))
	})

	It("should allow requests with valid token", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/user", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
	})
})

// var _ = Describe("APIKeyMiddleware", func() {
// 	const apiKey = "test-api-key"

// 	BeforeEach(func() {
// 		os.Setenv("API_KEY", apiKey)
// 	})

// 	AfterEach(func() {
// 		os.Unsetenv("API_KEY")
// 	})

// 	It("should reject requests without API key", func() {
// 		req, _ := http.NewRequest("GET", baseURL+"/health", nil)
// 		resp, err := http.DefaultClient.Do(req)
// 		Expect(err).To(BeNil())
// 		defer resp.Body.Close()
// 	})

// 	It("should reject requests with invalid API key", func() {
// 		req, _ := http.NewRequest("GET", baseURL+"/health", nil)
// 		req.Header.Set("X-API-KEY", "wrong-key")
// 		resp, err := http.DefaultClient.Do(req)
// 		Expect(err).To(BeNil())
// 		defer resp.Body.Close()
// 	})

// 	It("should allow requests with correct API key", func() {
// 		req, _ := http.NewRequest("GET", baseURL+"/health", nil)
// 		req.Header.Set("X-API-KEY", apiKey)
// 		resp, err := http.DefaultClient.Do(req)
// 		Expect(err).To(BeNil())
// 		defer resp.Body.Close()
// 	})
// })
