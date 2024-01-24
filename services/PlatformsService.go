package services

import (
	"gametracker/db"
	"gametracker/models"
	"gametracker/utils"
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

// CreatePlatform godoc
// @Tags platforms
// @Summary Create a platform
// @ID create-platform
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Platforms
// @Router /platforms [POST]
// @Param platform body models.Platforms true "Tag"
func CreatePlatform(c *gin.Context) {
	database := db.GetDatabase()
	var platform models.Platforms

	parsing := utils.CheckParse(c, &platform)
	if parsing == nil {
		return
	}
	
	checkPlatformExists := database.Where("name = ?", platform.Name).First(&platform)
	if checkPlatformExists.RowsAffected > 0 {
		c.JSON(400, gin.H{
			"error": "Platform already exists",
		})
		return
	}

	requestDb := database.Create(&platform)

	if requestDb.Error != nil {
		c.JSON(500, gin.H{
			"error": "Error creating platform",
		})
		return
	}

	c.JSON(200, gin.H{
		"platform": platform,
	})
}

// DeletePlatform godoc
// @Tags platforms
// @Summary Delete a platform
// @ID delete-platform
// @Produce  json
// @Success 200 {object} models.Platforms
// @Router /platforms/{id} [DELETE]
// @Param id path string true "Platform ID"
func DeletePlatform(c *gin.Context) {
	database := db.GetDatabase()
	var platform models.Platforms
	checkIfPlatformExists := database.Where("id = ?", c.Param("id")).First(&platform)
	if checkIfPlatformExists.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"error": "Platform does not exist",
		})
		return
	}

	requestDb := database.Delete(&platform)
	if requestDb.Error != nil {
		c.JSON(500, gin.H{
			"error": "Error deleting platform",
		})
		return
	}

	c.JSON(200, gin.H{
		"platform": platform,
	})
}
