package server_tests

import (
	"bytes"
	"encoding/json"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Run Handlers API", func() {
	var token string

	BeforeEach(func() {
		token = registerAndLogin("runuser@example.com", "password123", "runuser")
	})

	It("should upload, list, and delete a run", func() {
		// Upload a run
		runPayload := map[string]any{
			"route":    "LINESTRING(0 0,1 1)",
			"duration": "45",
			"distance": 10.5,
		}
		runBody, _ := json.Marshal(runPayload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/runs", bytes.NewReader(runBody))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		// List runs
		req, _ = http.NewRequest("GET", baseURL+"/protected/runs", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var runs []map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&runs)
		Expect(len(runs)).To(Equal(1))
		Expect(runs[0]["distance"]).To(Equal(10.5))
		Expect(runs[0]["duration"]).To(Equal("45"))
		runID := runs[0]["id"].(string)

		// Delete the run
		req, _ = http.NewRequest("DELETE", baseURL+"/protected/runs/"+runID, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		// List runs again (should be empty)
		req, _ = http.NewRequest("GET", baseURL+"/protected/runs", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		_ = json.NewDecoder(resp.Body).Decode(&runs)
		Expect(len(runs)).To(Equal(0))
	})

	It("should not allow unauthorized access", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/runs", nil)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(401))
	})

	It("should return error for invalid run ID on delete", func() {
		req, _ := http.NewRequest("DELETE", baseURL+"/protected/runs/invalid-uuid", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(400))
	})
})
