package services

import (
	"encoding/json"
	"gametracker/db"
	models "gametracker/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetTags godoc
// @Tags tags
// @Summary Get all tags
// @ID get-all-tags
// @Produce  json
// @Success 200 {object} models.Tags
// @Router /tags [get]
func GetTags(c *gin.Context) {
	database := db.GetDatabase()
	var tags []models.Tags

	requestDb := database.Find(&tags)

	if requestDb.Error != nil {
		log.Fatalf("Error getting tags: %v", requestDb.Error)
	}

	jsonTags, err := json.Marshal(tags)

	if err != nil {
		log.Fatalf("Error parsing tags: %v", err)
	}
	println(jsonTags)
	c.JSON(http.StatusOK, gin.H{
		"tags": tags,
	})
}

// CreateTag godoc
// @Tags tags
// @Summary Create a tag
// @ID create-tag
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Tags
// @Param tag body models.Tags true "Tag"
// @Router /tags [post]
func CreateTag(c *gin.Context) {
	database := db.GetDatabase()
	var tag models.Tags

	c.BindJSON(&tag)
	checkTagExists := database.Where("name = ?", tag.Name).First(&tag)
	if checkTagExists.RowsAffected > 0 { // tag already exists
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag already exists",
		})
		return
	}

	requestDb := database.Create(&tag)
	if requestDb.Error != nil { // error creating tag
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating tag",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"tag": tag,
	})
}

// DeleteTag godoc
// @Tags tags
// @Summary Delete a tag
// @ID delete-tag
// @Produce  json
// @Success 200 {object} models.Tags
// @Param id path int true "Tag ID"
// @Router /tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	database := db.GetDatabase()
	var tag models.Tags
	id := c.Params.ByName("id")

	requestDb := database.Where("id = ?", id).Delete(&tag)
	if requestDb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tag not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"tag": tag,
	})
}
