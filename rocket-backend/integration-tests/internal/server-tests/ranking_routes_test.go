package server_tests

import (
	"bytes"
	"encoding/json"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ranking Handlers API", func() {
	var tokenA string
	const userA = "rankA"
	const userB = "rankB"
	const userC = "rankC"

	BeforeEach(func() {
		tokenA = registerAndLogin("rankA@example.com", "password123", userA)
		_ = registerAndLogin("rankB@example.com", "password123", userB)
		_ = registerAndLogin("rankC@example.com", "password123", userC)

		// Set rocket points directly in DB for deterministic ranking
		_, err := testDbInstance.Exec("UPDATE users SET rocketpoints = 100 WHERE username = $1", userA)
		Expect(err).To(BeNil())
		_, err = testDbInstance.Exec("UPDATE users SET rocketpoints = 200 WHERE username = $1", userB)
		Expect(err).To(BeNil())
		_, err = testDbInstance.Exec("UPDATE users SET rocketpoints = 150 WHERE username = $1", userC)
		Expect(err).To(BeNil())

		// userA adds userB and userC as friends
		for _, friend := range []string{userB, userC} {
			payload := map[string]any{"friend_name": friend}
			body, _ := json.Marshal(payload)
			req, _ := http.NewRequest("POST", baseURL+"/protected/friends/add", bytes.NewReader(body))
			req.Header.Set("Authorization", "Bearer "+tokenA)
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			Expect(err).To(BeNil())
			resp.Body.Close()
		}
	})

	It("should return global user ranking", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/ranking/users", nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var ranking []map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&ranking)
		Expect(len(ranking)).To(BeNumerically(">=", 3))
		Expect(ranking[0]["rocket_points"]).To(Equal(float64(200)))
		Expect(ranking[1]["rocket_points"]).To(Equal(float64(150)))
		Expect(ranking[2]["rocket_points"]).To(Equal(float64(100)))
	})

	It("should return friends ranking for userA", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/ranking/friends", nil)
		req.Header.Set("Authorization", "Bearer "+tokenA)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var friendsRanking []map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&friendsRanking)
		Expect(len(friendsRanking)).To(Equal(3))
		Expect(friendsRanking[0]["username"]).To(Equal(userB))
		Expect(friendsRanking[1]["username"]).To(Equal(userC))
		Expect(friendsRanking[2]["username"]).To(Equal(userA))
	})
})
