package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	s := &Server{}
	r := gin.New()
	api := r.Group("/api/v1")
	{
		api.GET("/", s.HelloWorldHandler)
	}
	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/api/v1/", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Serve the HTTP request
	r.ServeHTTP(rr, req)
	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "Handler returned wrong status code")

	// Check the response body
	expected := "{\"message\":\"Hello World\"}"
	assert.JSONEq(t, expected, rr.Body.String(), "Handler returned unexpected body")
}
