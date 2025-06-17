package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	// r.Use(s.APIKeyMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-API-KEY"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	api := r.Group("/api/v1")
	{
		chatHub := NewChatHub()
		go chatHub.Run()

		api.GET("/health", s.HealthHandler)
		api.POST("/register", s.RegisterHandler)
		api.POST("/login", s.LoginHandler)
		api.POST("/logout", s.LogoutHandler)

		protected := api.Group("/protected")
		protected.Use(s.AuthMiddleware())
		{
			protected.GET("/", s.AuthenticatedHandler)
			protected.POST("/updateSteps", s.UpdateStepsHandler)

			//protected.POST("/settings/update", s.UpdateSettings)
			protected.GET("/settings", s.GetSettingsHandler)
			protected.POST("/settings/step-goal", s.UpdateStepGoalHandler)
			protected.POST("/settings/image", s.UpdateImageHandler)
			protected.DELETE("/settings/image", s.DeleteImageHandler)
			// protected.POST("/settings/userinfo", s.UpdateUserInfoHandler)

			protected.GET("/user", s.GetUserHandler)
			protected.DELETE("/user", s.DeleteUserHandler)
			protected.GET("/user/:name", s.GetUserByNameHandler)
			protected.POST("/user/statistics", s.GetUserStatisticsHandler)
			protected.POST("/user/image", s.GetUserImageHandler)
			protected.GET("/user/rocketpoints", s.GetRocketPointsHandler)
			protected.GET("/users", s.GetAllUsersHandler)

			protected.GET("/challenges/new", s.GetDailyChallengesHandler)
			protected.POST("/challenges/complete", s.CompleteChallengeHandler)
			protected.GET("/challenges/progress", s.GetDailyChallengeProgress)
			protected.POST("/challenges/invite", s.InviteFriendChallenge)

			protected.GET("/ranking/users", s.GetUserRankingHandler)
			protected.GET("/ranking/friends", s.GetFriendsRankedHandler)

			protected.GET("/friends", s.GetAllFriendsHandler)
			protected.POST("/friends/add", s.AddFriendHandler)
			protected.DELETE("/friends/:name", s.DeleteFriendHandler)
			protected.GET("/followers/:id", s.GetFollowersHandler)
			protected.GET("/following/:id", s.GetFollowingHandler)

			protected.POST("/runs", s.UploadRunHandler)
			protected.GET("/runs", s.GetAllRunsHandler)
			protected.DELETE("/runs/:id", s.DeleteRunHandler)
			protected.POST("/runs/plan", s.PlanRunHandler)
			protected.GET("/runs/plan", s.GetPlannedRunHandler)
			protected.DELETE("/runs/plan/:id", s.DeletePlannedRunHandler)

			protected.GET("/activites", s.GetActivityHandler)

			protected.GET("/ws/chat", s.ChatWebSocketHandler(chatHub))
			protected.GET("/chat/history", s.GetChatHistoryHandler)
		}
	}

	return r
}
