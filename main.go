package main

import (
	"log"
	"task5-pbi-btpns-raymondakkaseljayaimanuel/database"
	"task5-pbi-btpns-raymondakkaseljayaimanuel/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	// Connect to database
	database.Connect()
	// Run migrations
	database.Migrate()

	// Setup the router
	r := router.SetupRouter()

	// Run the server
	r.Run(":8080")
}
