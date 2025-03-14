package main

import (
	"user-microservice/database"
	"user-microservice/models"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
	println("Migrations completed successfully!")
}
