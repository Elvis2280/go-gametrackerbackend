package routes

import (
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupGamesRoutes(r *gin.Engine) {
	gamesRoutes := r.Group("/api/games")
	gamesRoutes.GET("", services.GetGames)
	gamesRoutes.POST("", services.CreateGame)
	gamesRoutes.DELETE("/:id", services.DeleteGame)
}
