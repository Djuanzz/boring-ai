package routes

import (
	"github.com/Djuanzz/boring-ai/controllers"

	"github.com/gin-gonic/gin"
)

func OpenAIRoutes(router *gin.Engine, openAIController controllers.OpenAIController) {
	api := router.Group("/api/ai")
	{
		api.POST("/chat", openAIController.GenerateResponse)
	}
}
