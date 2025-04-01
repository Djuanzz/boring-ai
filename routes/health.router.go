package routes

import (
	"github.com/Djuanzz/boring-ai/controllers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.Engine, hc controllers.HealthController) {
	routes := r.Group("/api/health")
	{
		routes.GET("/ping", hc.CheckPing)
	}
}
