package routes

import (
	"github.com/Djuanzz/boring-ai/controllers"
	"github.com/gin-gonic/gin"
)

func Search(r *gin.Engine, sc controllers.SearchController) {
	routes := r.Group("/api/task")
	{
		routes.POST("/search", sc.HandleSearch)
		routes.POST("/search/detail", sc.PlaceDetail)
	}
}
