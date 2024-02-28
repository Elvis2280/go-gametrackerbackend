package main

import (
	_ "gametracker/docs"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
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
	//db.ConnectDatabase()
	r := gin.Default()
	r.Use(gin.Logger())

	// Create a new Fizz instance from the Gin engine.
	f := fizz.NewFromEngine(r)

	// Add Open API description
	//infos := &openapi.Info{
	//	Title:       "Game tracker API",
	//	Description: "This is my Service API",
	//	Version:     "0.1",
	//}

	// Create an endpoint for openapi.json file
	f.GET("/openapi.json", nil, f.OpenAPI(infos, "json"))
	//// Now add a UI handler
	//r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//routes.SetupTagsRoutes(r)
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
