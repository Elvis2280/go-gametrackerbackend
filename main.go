package main

import (
	"gametracker/db"
	_ "gametracker/docs"
	"gametracker/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

// @title Game tracker API
// @description Game tracker app API
// @BasePath /api
// @TermsOfServiceUrl https://erudev.page
// @host      localhost:8080
// @BasePath  /api
// @Security Bearer
// @SecurityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	db.ConnectDatabase()
	r := gin.Default()
	r.Use(gin.Logger())

	f := fizz.NewFromEngine(r)

	// OpenAPI infos
	infos := &openapi.Info{
		Title:       "Game tracker API",
		Description: "This is my Service API",
		Version:     "0.1",
	}

	// open api routes
	f.GET("/openapi.json", nil, f.OpenAPI(infos, "json"))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// app routes
	routes.SetupTagsRoutes(r)
	//routes.SetupPlatformsRoutes(r)
	//routes.SetupGamesRoutes(r)
	//routes.SetupUsersRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()

	if err != nil {
		panic(err)
	}
}
