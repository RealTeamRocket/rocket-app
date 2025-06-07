package server_tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Public Handlers API", func() {
	It("should return health status", func() {
		resp, err := http.Get(baseURL + "/health")
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
	})

	It("should register a user", func() {
		payload := map[string]interface{}{
			"email":    "test@example.com",
			"password": "password123",
			"username": "testuser",
		}
		body, _ := json.Marshal(payload)
		resp, err := http.Post(baseURL+"/register", "application/json", bytes.NewReader(body))
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		b, _ := io.ReadAll(resp.Body)
		var result map[string]interface{}
		_ = json.Unmarshal(b, &result)
		Expect(result["message"]).To(Equal("User registered successfully"))
	})

	It("should not register a user with duplicate email", func() {
		payload := map[string]interface{}{
			"email":    "dupe@example.com",
			"password": "password123",
			"username": "dupeuser",
		}
		body, _ := json.Marshal(payload)
		resp, err := http.Post(baseURL+"/register", "application/json", bytes.NewReader(body))
		Expect(err).To(BeNil())
		resp.Body.Close()

		// Try to register again with same email
		resp2, err := http.Post(baseURL+"/register", "application/json", bytes.NewReader(body))
		Expect(err).To(BeNil())
		defer resp2.Body.Close()
		Expect(resp2.StatusCode).To(Equal(400))
		b, _ := io.ReadAll(resp2.Body)
		var result map[string]interface{}
		_ = json.Unmarshal(b, &result)
		Expect(result["error"]).To(Equal("Email already exists"))
	})

	It("should login a user", func() {
		// Register first
		payload := map[string]interface{}{
			"email":    "login@example.com",
			"password": "password123",
			"username": "loginuser",
		}
		body, _ := json.Marshal(payload)
		resp, err := http.Post(baseURL+"/register", "application/json", bytes.NewReader(body))
		Expect(err).To(BeNil())
		resp.Body.Close()

		// Login
		loginPayload := map[string]interface{}{
			"email":    "login@example.com",
			"password": "password123",
		}
		loginBody, _ := json.Marshal(loginPayload)
		resp, err = http.Post(baseURL+"/login", "application/json", bytes.NewReader(loginBody))
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		b, _ := io.ReadAll(resp.Body)
		var result map[string]interface{}
		_ = json.Unmarshal(b, &result)
		Expect(result).To(HaveKey("token"))
	})

	It("should not login with wrong password", func() {
		// Register first
		payload := map[string]interface{}{
			"email":    "wrongpass@example.com",
			"password": "password123",
			"username": "wrongpassuser",
		}
		body, _ := json.Marshal(payload)
		resp, err := http.Post(baseURL+"/register", "application/json", bytes.NewReader(body))
		Expect(err).To(BeNil())
		resp.Body.Close()

		// Login with wrong password
		loginPayload := map[string]interface{}{
			"email":    "wrongpass@example.com",
			"password": "wrongpassword",
		}
		loginBody, _ := json.Marshal(loginPayload)
		resp, err = http.Post(baseURL+"/login", "application/json", bytes.NewReader(loginBody))
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(401))
		b, _ := io.ReadAll(resp.Body)
		var result map[string]interface{}
		_ = json.Unmarshal(b, &result)
		Expect(result["error"]).To(Equal("Invalid username or password"))
	})

	It("should logout a user", func() {
		resp, err := http.Post(baseURL+"/logout", "application/json", nil)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		b, _ := io.ReadAll(resp.Body)
		var result map[string]interface{}
		_ = json.Unmarshal(b, &result)
		Expect(result["message"]).To(Equal("Logged out successfully"))
	})
})
