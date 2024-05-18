package routes

import (
	"gametracker/middleware"
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupGamesRoutes(r *gin.Engine) {
	gamesRoutes := r.Group("/api/games")
	gamesRoutes.Use(middleware.Auth())
	gamesRoutes.GET("", services.GetGames)
	gamesRoutes.GET("/count", services.GetCountGames)
	gamesRoutes.POST("", services.CreateGame)
	gamesRoutes.DELETE("/:id", services.DeleteGame)
	gamesRoutes.PUT("/:id", services.UpdateGame)
}
