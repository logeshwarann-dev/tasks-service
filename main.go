package main

import (
	"tasks-service/handlers"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize MongoDB connection
	handlers.ConnectDB()

	// Initialize the Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Allow cookies
	}))

	// Register the routes
	handlers.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
