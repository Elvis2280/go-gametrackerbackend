package services

import (
	"gametracker/db"
	"gametracker/models"
	"github.com/gin-gonic/gin"
)

// GetPlatforms godoc
// @Tags platforms
// @Summary Get all platforms
// @ID get-all-platforms
// @Produce  json
// @Success 200 {object} []models.Platforms
// @Router /platforms [get]
func GetPlatforms(c *gin.Context) {
	database := db.GetDatabase()
	var platforms []models.Platforms

	requestDb := database.Find(&platforms)
	if requestDb.Error != nil {
		c.JSON(500, gin.H{
			"error": "Error getting platforms",
		})
		return
	}

	c.JSON(200, gin.H{
		"platforms": platforms,
	})
}
