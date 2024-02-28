package routes

import (
	"gametracker/db"
	"gametracker/models"
	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(r *gin.Engine) {
	//r.POST("/api/signup", services.SignUp)
	//r.POST("/api/login", services.Login)
	r.GET("/api/users", func(c *gin.Context) {
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
	})
}
