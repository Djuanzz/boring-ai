package routes

import (
	"github.com/Djuanzz/boring-ai/controllers"
	"github.com/gin-gonic/gin"
)

func Scrape(r *gin.Engine, sc controllers.ScrapeController) {
	routes := r.Group("/api/task")
	{
		routes.POST("/scrape", sc.GetReviews)
	}
}
