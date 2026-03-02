package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"pietroballarin.com/paninup-backend/internal/database"
	"pietroballarin.com/paninup-backend/internal/model"
	"pietroballarin.com/paninup-backend/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Connect to database
	database.Connect()

	// Migrations
	database.DB.AutoMigrate(&model.User{})

	// Server initialization
	server := gin.Default()

	// Service setup
	user_service := service.NewUserService(database.DB)

	// Routes setup
	server.POST("/user/register", user_service.Create)

	// Server listening
	server.Run()
}
