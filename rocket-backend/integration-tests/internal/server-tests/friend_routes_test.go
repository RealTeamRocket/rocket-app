package server_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func registerAndLoginFriend(email, password, username string) string {
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
	var result map[string]any
	_ = json.NewDecoder(resp.Body).Decode(&result)
	token := result["token"].(string)
	return token
}

var _ = Describe("Friends API", func() {
	var tokenA string
	const userA = "userA"
	const userB = "userB"

	BeforeEach(func() {
		tokenA = registerAndLoginFriend("userA@example.com", "password123", userA)
		_ = registerAndLoginFriend("userB@example.com", "password123", userB)
	})

	It("should add, list, and delete friends", func() {
		// Add userB as friend to userA
		payload := map[string]any{"friend_name": userB}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/friends/add", bytes.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+tokenA)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		// List friends for userA
		req, _ = http.NewRequest("GET", baseURL+"/protected/friends", nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var friends []map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&friends)
		Expect(len(friends)).To(Equal(1))
		Expect(friends[0]["username"]).To(Equal(userB))

		// Delete userB as friend from userA
		req, _ = http.NewRequest("DELETE", fmt.Sprintf("%s/protected/friends/%s", baseURL, userB), nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		// List friends again for userA
		req, _ = http.NewRequest("GET", baseURL+"/protected/friends", nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
	})

	It("should not add a non-existent user as friend", func() {
		payload := map[string]any{"friend_name": "ghostuser"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/friends/add", bytes.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+tokenA)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(404))
	})

	It("should not delete a non-friend", func() {
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/protected/friends/%s", baseURL, userB), nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(500)) // Internal error or 404 depending on backend logic
	})

	It("should return correct followers and following lists", func() {
		// Register and login users
		tokenA := registerAndLoginFriend("followerA@example.com", "password123", "followerA")
		_ = registerAndLoginFriend("followedB@example.com", "password123", "followedB")

		// Add followedB as friend to followerA
		payload := map[string]any{"friend_name": "followedB"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/friends/add", bytes.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+tokenA)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		// Get user IDs for both users (simulate what backend expects)
		// We'll use the /protected/user endpoint to get the ID for followerA
		req, _ = http.NewRequest("GET", baseURL+"/protected/user", nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var userAInfo map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&userAInfo)
		userAID := userAInfo["id"].(string)

		// Now get userB's ID
		req, _ = http.NewRequest("GET", baseURL+"/protected/user/"+ "followedB", nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var userBInfo map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&userBInfo)
		userBID := userBInfo["id"].(string)

		// Check following for followerA (should include followedB)
		req, _ = http.NewRequest("GET", baseURL+"/protected/following/"+userAID, nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var following []map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&following)
		Expect(len(following)).To(Equal(1))
		Expect(following[0]["username"]).To(Equal("followedB"))

		// Check followers for followedB (should include followerA)
		req, _ = http.NewRequest("GET", baseURL+"/protected/followers/"+userBID, nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var followers []map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&followers)
		Expect(len(followers)).To(Equal(1))
		Expect(followers[0]["username"]).To(Equal("followerA"))
	})
})
