package routes

import (
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupGamesRoutes(r *gin.Engine) {
	r.GET("/api/games", services.GetGames)
	r.POST("/api/games", services.CreateGame)
	r.DELETE("/api/games/:id", services.DeleteGame)
}
