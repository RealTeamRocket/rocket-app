package main

import (
	"fmt"

	"github.com/RealTeamRocket/rocket-app/rocket-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize routes
	routes.RegisterRoutes(r)

	port := "8080"
	fmt.Println("Server is running on port", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
