package services

import (
	"gametracker/constants"
	"gametracker/db"
	"gametracker/models"
	"gametracker/utils"
	"github.com/gin-gonic/gin"
	"math"
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
	db := db.GetDatabase()
	games := []models.Game{}
	email := c.Query("email")
	limit := c.Query("limit")
	isActiveGames := c.Query("isActiveGames")
	page := c.Query("page")
	searchGame := c.Query("search")

	if limit == "" {
		limit = "10"
	}
	if isActiveGames == "" {
		isActiveGames = "true"
	}
	if page == "" {
		page = "1"
	}

	isActiveGamesParsed, err := strconv.ParseBool(isActiveGames)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid isActiveGames"})
		return

	}

	var totalItems int64

	limitParsed, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	pageParsed, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	if pageParsed < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	// Sort to the database
	if searchGame != "" {
		db = db.Where("LOWER(name) LIKE LOWER(?)", "%"+searchGame+"%")
	}

	if isActiveGamesParsed {
		db.Model(&models.Game{}).Where("email = ?", email).Not("status = ?", constants.Completed).Count(&totalItems)
	} else {
		db.Model(&models.Game{}).Where("email = ?", email).Where("status = ?", constants.Completed).Count(&totalItems)
	}

	var nextPage int
	roundedTotalItems := math.Ceil(float64(totalItems) / float64(limitParsed))
	if float64(pageParsed) > roundedTotalItems {
		if roundedTotalItems == 0 {
			nextPage = 1
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
			return
		}
	} else {
		nextPage = int(pageParsed + 1)
	}

	offset := (pageParsed - 1) * limitParsed

	if isActiveGamesParsed {
		db.Where("email = ?", email).Not("status = ?", constants.Completed).Offset(int(offset)).Limit(int(limitParsed)).Preload("Platforms").Preload("Tags").Find(&games)
	} else {
		db.Where("email = ?", email).Where("status = ?", constants.Completed).Offset(int(offset)).Limit(int(limitParsed)).Preload("Platforms").Preload("Tags").Find(&games)
	}

	//db.Offset(int(offset)).Limit(int(limitParsed)).Preload("Platforms").Preload("Tags").Find(&games)

	pagination := map[string]int{
		"totalPages":  int(totalItems / limitParsed),
		"currentPage": int(pageParsed),
		"nextPage":    nextPage,
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       games,
		"pagination": pagination,
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

func GetCountGames(c *gin.Context) {
	database := db.GetDatabase()
	email := strings.ToLower(c.Query("email")) // get email from query
	var activeCount int64
	var completedCount int64

	database.Model(&models.Game{}).Where("email = ? AND status = ?", email, constants.Completed).Count(&completedCount)
	database.Model(&models.Game{}).Where("email = ?", email).Not("status = ?", constants.Completed).Count(&activeCount)

	counts := map[string]int{
		"active":    int(activeCount),
		"completed": int(completedCount),
	}

	c.JSON(http.StatusOK, gin.H{
		"counts": counts,
	})
}

func calculatePagination(isFirstPage bool, hasPagination bool, limit int, games []models.Game, pointsNext bool) utils.PaginationInfo {
	pagination := utils.PaginationInfo{}
	nextCur := utils.Cursor{}
	prevCur := utils.Cursor{}
	if isFirstPage {
		if hasPagination {
			nextCur := utils.CreateCursor(strconv.Itoa(int(games[limit-1].ID)), games[limit-1].CreatedAt, true)
			pagination = utils.GeneratePager(nextCur, nil)
		}
	} else {
		if pointsNext {
			// if pointing next, it always has prev but it might not have next
			if hasPagination {
				nextCur = utils.CreateCursor(strconv.Itoa(int(games[limit-1].ID)), games[limit-1].CreatedAt, true)
			}
			prevCur = utils.CreateCursor(strconv.Itoa(int(games[0].ID)), games[0].CreatedAt, false)
			pagination = utils.GeneratePager(nextCur, prevCur)
		} else {
			// this is case of prev, there will always be nest, but prev needs to be calculated
			nextCur = utils.CreateCursor(strconv.Itoa(int(games[limit-1].ID)), games[limit-1].CreatedAt, true)
			if hasPagination {
				prevCur = utils.CreateCursor(strconv.Itoa(int(games[0].ID)), games[0].CreatedAt, false)
			}
			pagination = utils.GeneratePager(nextCur, prevCur)
		}
	}
	return pagination
}
