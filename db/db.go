package db

import (
	"gametracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var databaseUrlConnection string
var database *gorm.DB

func init() {
	println("Loading environment variables...")
	databaseUrlConnection = os.Getenv("DATABASE_URL")
}

func ConnectDatabase() {
	println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(databaseUrlConnection), &gorm.Config{}) // connect to database

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	errMigrate := db.AutoMigrate(&models.User{}, &models.Game{}, &models.Platforms{}, &models.Tags{}) // migrate models to database

	if errMigrate != nil {
		log.Fatal("Error migrating models")
	}

	pgDb, err := db.DB() // get database instance
	if err != nil {
		log.Fatal("Error getting database instance")
	}

	pgDb.SetMaxOpenConns(10)           // 10 connections
	pgDb.SetMaxIdleConns(10)           // 10 connections
	pgDb.SetConnMaxLifetime(time.Hour) // 1 hour
	database = db
	println("Database connected!")
}

func GetDatabase() (db *gorm.DB) {
	sqlDb, err := database.DB()

	if err != nil {
		log.Fatal("Error getting database instance")
	}

	if err := sqlDb.Ping(); err != nil {
		log.Fatal("Error pinging database")
	}

	return database

}
