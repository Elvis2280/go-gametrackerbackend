package routes

import (
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupPlatformsRoutes(r *gin.Engine) {
	r.GET("/api/platforms", services.GetPlatforms)
	r.POST("/api/platforms", services.CreatePlatform)
	r.DELETE("/api/platforms/:id", services.DeletePlatform)
}
