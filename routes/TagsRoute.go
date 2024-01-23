package routes

import (
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupTagsRoutes(r *gin.Engine) {
	r.GET("/api/tags", services.GetTags)
	r.POST("/api/tags", services.CreateTag)
	r.DELETE("/api/tags/:id", services.DeleteTag)
}
