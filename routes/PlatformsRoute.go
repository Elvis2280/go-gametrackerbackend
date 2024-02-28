package routes

import (
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupPlatformsRoutes(r *gin.Engine) {
	platformRoutes := r.Group("/api/platforms")
	//platformRoutes.Use(middleware.Auth())
	//platformRoutes.GET("", services.GetPlatforms)
	//platformRoutes.POST("", services.CreatePlatform)
	platformRoutes.DELETE("/:id", services.DeletePlatform)
}
