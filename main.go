package main

import (
	"fmt"

	"github.com/Djuanzz/boring-ai/config"
	"github.com/Djuanzz/boring-ai/controllers"
	"github.com/Djuanzz/boring-ai/middleware"
	"github.com/Djuanzz/boring-ai/routes"
	"github.com/Djuanzz/boring-ai/services"
	"github.com/gin-gonic/gin"
)

var (
	healthService services.HealthService = services.NewHealthService()

	healthController controllers.HealthController = controllers.NewHealthController(healthService)
)

func main() {
	fmt.Println("Boring AI")

	cfg := config.LoadEnv()

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	routes.Health(server, healthController)

	if err := server.Run(":" + cfg.Port); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err.Error())
	}

	fmt.Println("Server starting on port", cfg.Port)
}
