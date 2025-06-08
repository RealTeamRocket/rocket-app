package server_tests

import (
	"bytes"
	"encoding/json"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func registerAndLoginActivity(email, password, username string) string {
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

var _ = Describe("Activity Handlers API", func() {
	var token string

	BeforeEach(func() {
		token = registerAndLoginActivity("activity@example.com", "password123", "activityuser")
	})

	It("should get activities for the user (empty at first)", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/activites", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result).To(HaveKey("username"))
		Expect(result).To(HaveKey("activities"))
		var activities []any
		if result["activities"] != nil {
			activities = result["activities"].([]any)
		} else {
			activities = []any{}
		}
		Expect(len(activities)).To(Equal(0))
	})

	It("should create an activity in the DB by uploading a run", func() {
		runPayload := map[string]any{
			"route":    "LINESTRING(0 0,1 1)",
			"duration": "30",
			"distance": 5.0,
		}
		runBody, _ := json.Marshal(runPayload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/runs", bytes.NewReader(runBody))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		var count int
		err = testDB.DbInstance.QueryRow("SELECT COUNT(*) FROM activities").Scan(&count)
		Expect(err).To(BeNil())
		Expect(count).To(Equal(1))

		var message string
		err = testDB.DbInstance.QueryRow("SELECT message FROM activities ORDER BY time DESC LIMIT 1").Scan(&message)
		Expect(err).To(BeNil())
		Expect(message).To(ContainSubstring("Completed a 5.00 km run"))
	})

})
