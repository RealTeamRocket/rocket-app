package routes

import (
	"github.com/RealTeamRocket/rocket-app/rocket-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the API endpoints
func RegisterRoutes(r *gin.Engine) {
	r.GET("/hello", handlers.HelloHandler)
}
