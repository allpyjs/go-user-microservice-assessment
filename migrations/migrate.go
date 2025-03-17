package main

import (
	"log"

	"golang-microservice/config"
	"golang-microservice/models"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migration completed successfully")
}
