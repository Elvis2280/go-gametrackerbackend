package main

import (
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
	//db.ConnectDatabase()
	r := gin.Default()
	r.Use(gin.Logger())

	// Create a new Fizz instance from the Gin engine.
	f := fizz.NewFromEngine(r)

	// Add Open API description
	infos := &openapi.Info{
		Title:       "Game tracker API",
		Description: "This is my Service API",
		Version:     "0.1",
	}

	// Create an endpoint for openapi.json file
	f.GET("/openapi.json", nil, f.OpenAPI(infos, "json"))
	//// Now add a UI handler
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.SetupTagsRoutes(r)
	routes.SetupPlatformsRoutes(r)
	routes.SetupGamesRoutes(r)
	routes.SetupUsersRoutes(r)
	err := r.Run()

	if err != nil {
		panic(err)
	}
}
