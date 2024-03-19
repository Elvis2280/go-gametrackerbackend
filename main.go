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
	"os"
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

	isDev := os.Getenv("IS_DEV")
	var corsAllowed string
	if isDev == "true" {
		corsAllowed = "http://localhost:3000"
	} else {
		corsAllowed = "https://gametracker-elvisdev.netlify.app"
	}

	// open api routes
	f.GET("/openapi.json", nil, f.OpenAPI(infos, "json"))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Cors
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", corsAllowed)
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "https://gametracker-elvisdev.netlify.app")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	// app routes
	routes.SetupTagsRoutes(r)
	routes.SetupPlatformsRoutes(r)
	routes.SetupGamesRoutes(r)
	routes.SetupUsersRoutes(r)
	err := r.Run()

	if err != nil {
		panic(err)
	}
}
