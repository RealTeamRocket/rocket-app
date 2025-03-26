package routes

import (
	"github.com/RealTeamRocket/rocket-app/rocket-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

// registerAPIRoutes registers all routes within the given router group
func registerAPIRoutes(api *gin.RouterGroup) {
	api.GET("/hello", handlers.HelloHandler)
}

// RegisterRoutes sets up the API endpoints with the /api prefix
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	registerAPIRoutes(api)
}
