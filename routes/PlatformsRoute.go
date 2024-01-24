package routes

import (
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupPlatformsRoutes(r *gin.Engine) {
	r.GET("/api/platforms", services.GetPlatforms)
}
