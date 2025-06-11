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

	It("should create, list, and delete a planned run", func() {
		// Create a planned run
		planPayload := map[string]any{
			"route":    "LINESTRING(2 2,3 3)",
			"name":     "Morning Run",
			"distance": 5.0,
		}
		planBody, _ := json.Marshal(planPayload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/runs/plan", bytes.NewReader(planBody))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		// List planned runs
		req, _ = http.NewRequest("GET", baseURL+"/protected/runs/plan", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var plannedRuns []map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&plannedRuns)
		Expect(len(plannedRuns)).To(Equal(1))
		Expect(plannedRuns[0]["name"]).To(Equal("Morning Run"))
		Expect(plannedRuns[0]["distance"]).To(Equal(5.0))
		plannedRunID := plannedRuns[0]["id"].(string)

		// Delete the planned run
		req, _ = http.NewRequest("DELETE", baseURL+"/protected/runs/plan/"+plannedRunID, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		// List planned runs again (should be empty)
		req, _ = http.NewRequest("GET", baseURL+"/protected/runs/plan", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		_ = json.NewDecoder(resp.Body).Decode(&plannedRuns)
		Expect(len(plannedRuns)).To(Equal(0))
	})

	It("should not allow duplicate planned run names", func() {
		planPayload := map[string]any{
			"route":    "LINESTRING(4 4,5 5)",
			"name":     "Duplicate Run",
			"distance": 7.0,
		}
		planBody, _ := json.Marshal(planPayload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/runs/plan", bytes.NewReader(planBody))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		// Try to create another planned run with the same name
		req, _ = http.NewRequest("POST", baseURL+"/protected/runs/plan", bytes.NewReader(planBody))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(400))
	})

	It("should not allow unauthorized access to planned runs", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/runs/plan", nil)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(401))
	})

	It("should return error for invalid planned run ID on delete", func() {
		req, _ := http.NewRequest("DELETE", baseURL+"/protected/runs/plan/invalid-uuid", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(400))
	})
})
