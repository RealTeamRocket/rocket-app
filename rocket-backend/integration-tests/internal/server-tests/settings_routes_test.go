package server_tests

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Settings Handlers API", func() {
	var token string

	BeforeEach(func() {
		token = registerAndLogin("settingsuser@example.com", "password123", "settingsuser")
	})

	It("should get user settings", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/settings", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var settings map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&settings)
		Expect(settings).To(HaveKey("step_goal"))
	})

	It("should update step goal", func() {
		payload := map[string]any{"stepGoal": 12345}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/settings/step-goal", bytes.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result["message"]).To(Equal("Step goal updated successfully"))

		// Confirm update
		req, _ = http.NewRequest("GET", baseURL+"/protected/settings", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var settings map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&settings)
		Expect(settings["step_goal"]).To(Equal(float64(12345)))
	})

	It("should reject invalid step goal", func() {
		payload := map[string]any{"stepGoal": 0}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/settings/step-goal", bytes.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(400))
	})

	It("should update user image", func() {
		// Prepare a small image file (use any small file for test)
		imgPath := filepath.Join(os.TempDir(), "testimg.png")
		os.WriteFile(imgPath, []byte{0x89, 0x50, 0x4E, 0x47}, 0644) // PNG header bytes

		file, _ := os.Open(imgPath)
		defer file.Close()

		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		fw, _ := w.CreateFormFile("image", "testimg.png")
		file.Seek(0, 0)
		_, _ = file.WriteTo(fw)
		w.Close()

		req, _ := http.NewRequest("POST", baseURL+"/protected/settings/image", &b)
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", w.FormDataContentType())
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result["message"]).To(Equal("Image updated successfully"))
	})

	It("should delete user image", func() {
		// First, upload an image so there is something to delete
		imgPath := filepath.Join(os.TempDir(), "testimg_del.png")
		os.WriteFile(imgPath, []byte{0x89, 0x50, 0x4E, 0x47}, 0644)
		file, _ := os.Open(imgPath)
		defer file.Close()

		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		fw, _ := w.CreateFormFile("image", "testimg_del.png")
		file.Seek(0, 0)
		_, _ = file.WriteTo(fw)
		w.Close()

		req, _ := http.NewRequest("POST", baseURL+"/protected/settings/image", &b)
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", w.FormDataContentType())
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		resp.Body.Close()

		// Now, delete the image
		req, _ = http.NewRequest("DELETE", baseURL+"/protected/settings/image", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err = http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result["message"]).To(Equal("Image deleted successfully"))
	})

	It("should update user info (name and email)", func() {
		payload := map[string]any{
			"name":  "New Name",
			"email": "newemail@example.com",
		}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/settings/userinfo", bytes.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result["message"]).To(Equal("User info updated successfully"))
	})

	It("should update user password with correct current password", func() {
		payload := map[string]any{
			"currentPassword": "password123",
			"newPassword":     "newpass456",
		}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/settings/userinfo", bytes.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result["message"]).To(Equal("User info updated successfully"))
	})

	It("should reject password update with wrong current password", func() {
		payload := map[string]any{
			"currentPassword": "wrongpassword",
			"newPassword":     "newpass456",
		}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", baseURL+"/protected/settings/userinfo", bytes.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(401))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result["error"]).To(Equal("Current password incorrect"))
	})
})
