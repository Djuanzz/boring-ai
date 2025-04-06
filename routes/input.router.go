package routes

import (
	"github.com/Djuanzz/boring-ai/controllers"
	"github.com/gin-gonic/gin"
)

func Input(r *gin.Engine, tc controllers.InputController) {
	routes := r.Group("/api/task")
	{
		routes.POST("/input", tc.HandleInput)
	}
}
