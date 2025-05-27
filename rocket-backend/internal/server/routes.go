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
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
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
			protected.GET("/", s.AuthenticatedHandler)
			protected.POST("/updateSteps", s.UpdateStepsHandler)

			//protected.POST("/settings/update", s.UpdateSettings)
			protected.GET("/settings", s.GetSettingsHandler)
			protected.POST("/settings/step-goal", s.UpdateStepGoalHandler)
			protected.POST("/settings/image", s.UpdateImageHandler)

			protected.POST("/user/statistics", s.GetUserStatisticsHandler)
			protected.POST("/user/image", s.GetUserImageHandler)
			protected.GET("/user/rocketpoints", s.GetRocketPointsHandler)
			protected.GET("/users", s.GetAllUsersHandler)

			protected.GET("/challenges/new", s.GetDailyChallengesHandler)
			protected.POST("/challenges/complete", s.CompleteChallengeHandler)

			protected.GET("/ranking/users", s.GetUserRankingHandler)
			protected.GET("/ranking/friends", s.GetFriendsRankedHandler)

			protected.GET("/friends", s.GetAllFriendsHandler)
			protected.POST("/friends/add", s.AddFriendHandler)
			protected.DELETE("/friends/delete", s.DeleteFriendHandler)

			protected.POST("/runs", s.UploadRunHandler)
			protected.GET("/runs", s.GetAllRunsHandler)
		}
	}

	return r
}
