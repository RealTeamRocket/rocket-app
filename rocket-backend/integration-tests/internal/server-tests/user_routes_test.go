package server_tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func registerAndLogin(email, password, username string) string {
	// Register
	regPayload := map[string]any{
		"email":    email,
		"password": password,
		"username": username,
	}
	regBody, _ := json.Marshal(regPayload)
	http.Post(baseURL+"/register", "application/json", bytes.NewReader(regBody))

	// Login
	loginPayload := map[string]any{
		"email":    email,
		"password": password,
	}
	loginBody, _ := json.Marshal(loginPayload)
	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewReader(loginBody))
	Expect(err).To(BeNil())
	defer resp.Body.Close()
	Expect(resp.StatusCode).To(Equal(200))
	b, _ := io.ReadAll(resp.Body)
	var result map[string]any
	_ = json.Unmarshal(b, &result)
	token := result["token"].(string)
	return token
}

var _ = Describe("Protected Handlers API", func() {
	var token string

	BeforeEach(func() {
		token = registerAndLogin("protected@example.com", "password123", "protecteduser")
	})

	It("should get user statistics", func() {
		// Get user ID from /protected/user
		req, _ := http.NewRequest("GET", baseURL+"/protected/user", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var user map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&user)
		userID := user["id"].(string)

		// POST /protected/user/statistics
		payload := map[string]any{"id": userID}
		body, _ := json.Marshal(payload)
		req2, _ := http.NewRequest("POST", baseURL+"/protected/user/statistics", bytes.NewReader(body))
		req2.Header.Set("Authorization", "Bearer "+token)
		req2.Header.Set("Content-Type", "application/json")
		resp2, err := http.DefaultClient.Do(req2)
		Expect(err).To(BeNil())
		defer resp2.Body.Close()
		Expect(resp2.StatusCode).To(Equal(200))
	})

	It("should get user image (default empty)", func() {
		// Get user ID from /protected/user
		req, _ := http.NewRequest("GET", baseURL+"/protected/user", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var user map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&user)
		userID := user["id"].(string)

		// POST /protected/user/image
		payload := map[string]any{"user_id": userID}
		body, _ := json.Marshal(payload)
		req2, _ := http.NewRequest("POST", baseURL+"/protected/user/image", bytes.NewReader(body))
		req2.Header.Set("Authorization", "Bearer "+token)
		req2.Header.Set("Content-Type", "application/json")
		resp2, err := http.DefaultClient.Do(req2)
		Expect(err).To(BeNil())
		defer resp2.Body.Close()
		Expect(resp2.StatusCode).To(Equal(200))
		var imgResp map[string]any
		_ = json.NewDecoder(resp2.Body).Decode(&imgResp)
		Expect(imgResp).To(HaveKey("username"))
	})

	It("should get rocket points", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/user/rocketpoints", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var rp map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&rp)
		Expect(rp).To(HaveKey("rocket_points"))
	})

	It("should get all users", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/users", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var users []map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&users)
		Expect(users).To(BeNil())
	})
})
