package main

import (
	"log"
	"os"

	"golang-microservice/config"
	"golang-microservice/middleware"
	"golang-microservice/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Database
	config.ConnectDatabase()
	middleware.InitOAuth()

	// Set up router
	router := gin.Default()
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
