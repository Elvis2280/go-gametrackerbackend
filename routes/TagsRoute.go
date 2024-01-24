package routes

import (
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupTagsRoutes(r *gin.Engine) {
	tagsRoutes := r.Group("/api/tags")
	tagsRoutes.GET("", services.GetTags)
	tagsRoutes.POST("", services.CreateTag)
	tagsRoutes.DELETE("/:id", services.DeleteTag)
}
