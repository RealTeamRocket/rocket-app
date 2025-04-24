package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(s.APIKeyMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-API-KEY"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	api := r.Group("/api/v1")
	{
		api.GET("/health", s.HealthHandler)
		api.POST("/register", s.RegisterHandler)
		api.POST("/login", s.LoginHandler)

		protected := api.Group("/protected")
		protected.Use(s.AuthMiddleware())
		{
			protected.GET("/", s.Authenticated)
			protected.POST("/updateSteps", s.UpdateSteps)

			protected.POST("/settings/update", s.UpdateSettings)
			protected.GET("/settings", s.GetSettings)
			protected.POST("/image", s.GetUserImage)
		}
	}

	return r
}
