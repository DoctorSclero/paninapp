package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"pietroballarin.com/paninup-backend/internal/database"
	"pietroballarin.com/paninup-backend/internal/user"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Connect to database
	database.Connect()

	// Migrations
	database.DB.AutoMigrate(&user.User{})

	// Server initialization
	server := gin.Default()

	// Service setup
	userService := user.NewService(database.DB)

	// Routes setup
	server.POST("/users/register", userService.Register)
	// Server listening
	server.Run()
}
