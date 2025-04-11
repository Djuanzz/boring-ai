package routes

import (
	"github.com/Djuanzz/boring-ai/config"
	"github.com/Djuanzz/boring-ai/controllers"
	"github.com/Djuanzz/boring-ai/docs"
	"github.com/Djuanzz/boring-ai/middleware"
	"github.com/Djuanzz/boring-ai/services"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(cfg config.Config) *gin.Engine {
	docs.SwaggerInfo.Title = "Boring AI API"
	docs.SwaggerInfo.Description = "This is api documentation for boring ai."
	docs.SwaggerInfo.Version = "1.16.4"
	docs.SwaggerInfo.Host = "localhost:5000/api"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

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

	Health(server, healthController)
	Input(server, inputController)
	OpenAIRoutes(server, openAIController)
	Search(server, searchController)
	Scrape(server, scrapeController)

	return server
}
