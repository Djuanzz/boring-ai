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

func main() {
	fmt.Println("Boring AI")

	cfg := config.LoadEnv()

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	mapsClient := config.NewMapsClient(cfg.GMapsKey)
	openAIClient := config.NewOpenAIClient(cfg.OpenAIKey)

	healthService := services.NewHealthService()
	inputService := services.NewInputService()
	openAIService := services.NewOpenAIService(openAIClient)
	searchService := services.NewSearchService(mapsClient)
	scrapeService := services.NewScrapeService(cfg.SearchKey)

	healthController := controllers.NewHealthController(healthService)
	inputController := controllers.NewInputController(inputService)
	openAIController := controllers.NewOpenAIController(openAIService)
	searchController := controllers.NewSearchController(searchService)
	scrapeController := controllers.NewScrapeController(scrapeService)

	routes.Health(server, healthController)
	routes.Input(server, inputController)
	routes.OpenAIRoutes(server, openAIController)
	routes.Search(server, searchController)
	routes.Scrape(server, scrapeController)

	if err := server.Run(":" + cfg.Port); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err.Error())
	}

	fmt.Println("Server starting on port", cfg.Port)
}
