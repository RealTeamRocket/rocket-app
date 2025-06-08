package server_tests

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chat Handlers API", func() {
	var token string

	BeforeEach(func() {
		token = registerAndLogin("chathandler@example.com", "password123", "chathandleruser")
	})

	It("should return empty chat history at first", func() {
		req, _ := http.NewRequest("GET", baseURL+"/protected/chat/history", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result).To(HaveKey("messages"))
		var messages []any
		if result["messages"] != nil {
			messages = result["messages"].([]any)
		} else {
			messages = []any{}
		}
		Expect(len(messages)).To(Equal(0))
	})

	It("should show sent chat message in chat history", func() {
		httpURL := baseURL + "/protected/ws/chat"
		u, err := url.Parse(httpURL)
		Expect(err).To(BeNil())
		u.Scheme = "ws"
		header := http.Header{}
		header.Set("Authorization", "Bearer "+token)
		ws, _, err := websocket.DefaultDialer.Dial(u.String(), header)
		Expect(err).To(BeNil())
		defer ws.Close()

		msg := map[string]string{"message": "history test message"}
		msgBytes, _ := json.Marshal(msg)
		err = ws.WriteMessage(websocket.TextMessage, msgBytes)
		Expect(err).To(BeNil())

		// Wait for the message to be processed and stored
		time.Sleep(300 * time.Millisecond)

		req, _ := http.NewRequest("GET", baseURL+"/protected/chat/history", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		var result map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&result)
		Expect(result).To(HaveKey("messages"))
		messages := result["messages"].([]any)
		Expect(len(messages)).To(Equal(1))
		msgObj := messages[0].(map[string]any)
		Expect(msgObj["message"]).To(Equal("history test message"))
		Expect(msgObj["username"]).To(Equal("You"))
	})
})
