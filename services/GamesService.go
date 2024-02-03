package services

import (
	"gametracker/db"
	"gametracker/models"
	"gametracker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetGames godoc
// @Tags games
// @Summary Get all Games
// @ID get-all-games
// @Produce  json
// @Success 200 {object} []models.Game
// @Router /games [get]
func GetGames(c *gin.Context) {
	database := db.GetDatabase() // get database connection
	var games []models.Game      // create empty array of games

	requestDb := database.Find(&games) // get all games from database

	if requestDb.Error != nil { // 500
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error getting games",
		})
	}

	c.JSON(http.StatusOK, gin.H{ // 200
		"games": games,
	})
}

// CreateGame godoc
// @Tags games
// @Summary Create a game
// @ID create-game
// @Accept  json
// @Produce  json
// @Success 200 {object} models.CreateGame
// @Router /games [POST]
// @Param platform body models.CreateGame true "Tag"
func CreateGame(c *gin.Context) {
	database := db.GetDatabase()
	var game models.Game

	parsing := utils.CheckParse(c, &game)
	if parsing == nil {
		return
	}

	checkGameExists := database.Where("name = ?", game.Name).First(&game) // check if game already exists

	if checkGameExists.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Game already exists",
		})
		return
	}

	var platformData models.Platforms
	for _, platform := range game.Platforms { // add platforms to game
		database.FirstOrCreate(&platformData, models.Platforms{Name: platform.Name, IconName: platform.IconName})
		game.Platforms = append(game.Platforms, platformData)
	}
	game.Platforms = []models.Platforms{platformData}

	var tagData models.Tags
	for _, tag := range game.Tags { // add tags to game
		database.FirstOrCreate(&tagData, models.Tags{Name: tag.Name})
		game.Tags = append(game.Tags, tagData)
	}
	game.Tags = []models.Tags{tagData}

	requestDb := database.Create(&game) // create game

	if requestDb.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating game",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})
}

// DeleteGame godoc
// @Tags games
// @Summary Delete a game
// @ID delete-game
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Game
// @Router /games [DELETE]
// @Param platform body models.Game true "Tag"
func DeleteGame(c *gin.Context) {
	database := db.GetDatabase()
	var game models.Game

	checkIfGameExist := database.Where("id = ?", c.Param("id")).First(&game) // check if game exists
	if checkIfGameExist.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Game does not exist",
		})
		return
	}

	requestDb := database.Delete(&game)
	if requestDb.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting game",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})
}

// UpdateGame godoc
// @Tags games
// @Summary Update a game
// @ID update-game
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Game
// @Router /games/{id} [PUT]
// @Param platform body models.Game true "Tag"
func UpdateGame(c *gin.Context) {
	database := db.GetDatabase()
	var game models.Game

	checkIfGameExist := database.Where("id = ?", c.Param("id")).First(&game) // check if game exists
	if checkIfGameExist.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Game does not exist",
		})
		return
	}

	parsing := utils.CheckParse(c, &game)
	if parsing == nil {
		return
	}

	requestDb := database.Save(&game)
	if requestDb.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating game",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})

}
