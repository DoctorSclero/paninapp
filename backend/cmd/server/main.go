package main

import (
	"github.com/gin-gonic/gin"
	"pietroballarin.com/paninup-backend/internal/controllers"
	"pietroballarin.com/paninup-backend/internal/db"
	"pietroballarin.com/paninup-backend/internal/repositories"
)

func main() {
	// Database initialization
	db.InitDB()
	defer db.CloseDB()

	// Repository initialization
	userRepo := repositories.User{DB: db.DB}

	// Controller initialization
	userController := controllers.User{Repo: &userRepo}

	// Server initialization
	server := gin.Default()

	server.POST("/users/register", userController.Register)

	server.Run()
}
