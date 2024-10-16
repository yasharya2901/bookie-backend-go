package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yasharya2901/bookie-backend-go/config"
	"github.com/yasharya2901/bookie-backend-go/handlers"
	"github.com/yasharya2901/bookie-backend-go/models"
	"github.com/yasharya2901/bookie-backend-go/services"
)

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		host, user, password, dbname, dbport, sslmode, timezone,
	)

	db, err := config.InitDB(dsn)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.UserLocation{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	userLocationService := services.NewUserLocationService(db)

	userLocationHandler := handlers.NewUserLocationService(userLocationService)

	router := gin.Default()

	router.POST("/userlocation", userLocationHandler.CreateUserLocationHandler)
	router.GET("/userlocation/:appwrite_user_id", userLocationHandler.GetUserLocationHandler)
	router.POST("/appwrite/user/location", userLocationHandler.CreateUserFromAppwrite)

	if err := router.Run(port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
