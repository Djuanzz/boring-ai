package main

import (
	"fmt"
	"log"

	"github.com/Djuanzz/boring-ai/config"
	"github.com/Djuanzz/boring-ai/controllers"
	"github.com/Djuanzz/boring-ai/middleware"
	"github.com/Djuanzz/boring-ai/routes"
	"github.com/Djuanzz/boring-ai/services"
	"github.com/gin-gonic/gin"
	"googlemaps.github.io/maps"
)

var (
// healthService services.HealthService = services.NewHealthService()

// healthController controllers.HealthController = controllers.NewHealthController(healthService)
)

func main() {
	fmt.Println("Boring AI")

	cfg := config.LoadEnv()

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	client, err := maps.NewClient(maps.WithAPIKey(cfg.GMapsKey))
	if err != nil {
		log.Fatalf("Failed to create maps client: %v", err)
		panic(err)
	}

	healthService := services.NewHealthService()
	openAIService := services.NewOpenAIService(cfg)
	inputService := services.NewInputService()
	searchService := services.NewSearchService(client)
	scrapeService := services.NewScrapeService(cfg.SearchKey)

	healthController := controllers.NewHealthController(healthService)
	openAIController := controllers.NewOpenAIController(openAIService)
	inputController := controllers.NewInputController(inputService)
	searchController := controllers.NewSearchController(searchService)
	scrapeController := controllers.NewScrapeController(scrapeService)

	routes.Health(server, healthController)
	routes.OpenAIRoutes(server, openAIController)
	routes.Input(server, inputController)
	routes.Search(server, searchController)
	routes.Scrape(server, scrapeController)

	if err := server.Run(":" + cfg.Port); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err.Error())
	}

	fmt.Println("Server starting on port", cfg.Port)
}
