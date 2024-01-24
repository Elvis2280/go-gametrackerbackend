package routes

import (
	"gametracker/services"
	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(r *gin.Engine) {
	r.POST("/api/signup", services.SignUp)
	r.POST("/api/login", services.Login)
}
