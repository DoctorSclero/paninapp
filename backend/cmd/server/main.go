package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"pietroballarin.com/paninup-backend/internal/database"
	"pietroballarin.com/paninup-backend/internal/middleware"
	"pietroballarin.com/paninup-backend/internal/order"
	"pietroballarin.com/paninup-backend/internal/sandwich"
	"pietroballarin.com/paninup-backend/internal/types"
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
	database.DB.AutoMigrate(&user.User{}, &sandwich.Sandwich{}, &order.Order{}, &order.OrderItem{})

	// Server initialization
	server := gin.Default()

	// CORS Middleware
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Service setup
	userService := user.NewService(database.DB)
	orderService := order.NewService(database.DB)
	sandwichService := sandwich.NewService(database.DB)

	// Routes setup
	server.POST("/users/register", userService.Register)
	server.POST("/users/login", userService.Login)

	// Protected routes
	authorized := server.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/orders", orderService.Create)
		authorized.GET("/orders", orderService.List)
		authorized.GET("/sandwiches", sandwichService.List)

		// Manager only routes
		managers := authorized.Group("/")
		managers.Use(middleware.RequireRole(types.RoleManager))
		{
			managers.POST("/sandwiches", sandwichService.Create)
		}
	}

	// Server listening
	server.Run()
}
