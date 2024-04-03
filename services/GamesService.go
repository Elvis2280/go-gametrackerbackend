package services

import (
	"gametracker/db"
	"gametracker/models"
	"gametracker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
	var games []models.Game

	// create empty array of games
	email := strings.ToLower(c.Query("email")) // get email from query

	// get query params for pagination
	sortKey := c.Query("sortKey")
	status := c.Query("status")
	//tags := c.Query("tags")

	isAscendingStr := c.Query("isAscending")
	isAscending, err := strconv.ParseBool(isAscendingStr)
	if err != nil {
		isAscending = false
	}
	itemPerPageStr := c.Query("itemPerPage")
	itemsPerPage, err := strconv.Atoi(itemPerPageStr)
	if err != nil || itemsPerPage < 1 {
		itemsPerPage = 1
	}
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pag := &utils.Pagination{
		ItemsPerPage: itemsPerPage,
		CurrentPage:  page,
		SortKey:      sortKey,
		IsAscending:  isAscending,
		Status:       status,
		//Tags:         strings.Split(tags, ","),
	}

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email param is required",
		})
		return
	}
	var totalItemsInt int64

	database.Model(&models.Game{}).Where("Email = ?", email).Count(&totalItemsInt) // get total items for pagination

	database, paginatedData := utils.HandlePagination(database, pag, totalItemsInt)

	requestDb := database.Preload("Platforms").Preload("Tags").Where("Email = ?", email).Find(&games) // get all ga // mes with platforms and tags with preload

	if requestDb.Error != nil { // 500
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error getting games",
		})
	}

	c.JSON(http.StatusOK, gin.H{ // 200
		"games":      games,
		"pagination": paginatedData,
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

	checkGameExists := database.Where("Name = ?", game.Name).Find(&game) // check if game already exists

	if checkGameExists.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Game already exists",
		})
		return
	}

	var platformData []models.Platforms
	for _, platform := range game.Platforms { // add platforms to game
		var platformCreated models.Platforms
		database.FirstOrCreate(&platformCreated, models.Platforms{Name: platform.Name, IconName: platform.IconName})
		platformData = append(platformData, platformCreated)
	}
	game.Platforms = platformData // necessary to append platforms to game

	var tagData []models.Tags
	for _, tag := range game.Tags { // check if tag exists assign it to tag data if not create it
		var tagCreated models.Tags
		database.FirstOrCreate(&tagCreated, models.Tags{Name: tag.Name})
		tagData = append(tagData, tagCreated)
	}
	game.Tags = tagData // necessary to append tags to game

	requestDb := database.Create(&game) // create game

	database.Model(&game).Association("Platforms").Append(game.Platforms)
	database.Model(&game).Association("Tags").Append(game.Tags)

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

	var gameId = c.Param("id")
	if gameId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Game id is required",
		})
		return
	}

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
	var gameId = c.Param("id")

	checkIfGameExist := database.Where("id = ?", gameId).First(&game) // check if game exists
	if checkIfGameExist.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Game does not exist",
		})
		return
	}

	database.Model(&game).Association("Platforms").Clear()
	database.Model(&game).Association("Tags").Clear()
	// Delete the associations to avoid duplicates or keep the same associations

	parsing := utils.CheckParse(c, &game)
	if parsing == nil {
		return
	}

	var platformData []models.Platforms
	for _, platform := range game.Platforms { // add platforms to game
		var platformCreated models.Platforms
		database.FirstOrCreate(&platformCreated, models.Platforms{Name: platform.Name, IconName: platform.IconName})
		platformData = append(platformData, platformCreated)
	}
	game.Platforms = platformData // necessary to append platforms to game

	var tagData []models.Tags
	for _, tag := range game.Tags { // check if tag exists assign it to tag data if not create it
		var tagCreated models.Tags
		database.FirstOrCreate(&tagCreated, models.Tags{Name: tag.Name})
		tagData = append(tagData, tagCreated)
	}
	game.Tags = tagData // necessary to append tags to game

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
