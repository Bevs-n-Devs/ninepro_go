package main

import (
	"log"
	"os"

	"ninepro_go/pkg/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Check if the environment is production
	if os.Getenv("ENV") != "production" {
		// Load environment variables from .env in non-production environments
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file: %v", err) // Use Printf instead of Fatalf to continue even if the .env file is missing
		} else {
			log.Println(".env file loaded successfully")
		}
	} else {
		log.Println("Running in production mode, skipping .env file loading")
	}

	// Initialize Gin
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Define routes
	r.GET("/", handler.Index)
	r.GET("/home", handler.Home)
	r.POST("/contact", handler.Contact)

	// Run the server on port 8080
	r.Run(":8080")
}
