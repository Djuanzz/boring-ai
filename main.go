package main

import (
	"fmt"

	"github.com/Djuanzz/boring-ai/config"
	"github.com/Djuanzz/boring-ai/controllers"
	"github.com/Djuanzz/boring-ai/docs"
	"github.com/Djuanzz/boring-ai/middleware"
	"github.com/Djuanzz/boring-ai/routes"
	"github.com/Djuanzz/boring-ai/services"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	fmt.Println("Boring AI")

	cfg := config.LoadEnv()

	docs.SwaggerInfo.Title = "Boring AI API"
	docs.SwaggerInfo.Description = "This is api documentation for boring ai."
	docs.SwaggerInfo.Version = "1.16.4"
	docs.SwaggerInfo.Host = "localhost:5000/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
