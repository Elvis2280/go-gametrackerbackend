package routes

import (
	"gametracker/middleware"
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupTagsRoutes(r *gin.Engine) {
	tagsRoutes := r.Group("/api/tags")
	tagsRoutes.Use(middleware.Auth())
	tagsRoutes.GET("", services.GetTags)
	tagsRoutes.POST("", services.CreateTag)
	tagsRoutes.DELETE("/:id", services.DeleteTag)
}
