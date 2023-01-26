package main

import (
	"fmt"
	handler "thirdweb-go/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")

	// If .env file is not present
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Initialize gin router
	router := gin.Default()

	// Route handlers / endpoints
	router.POST("/generate", handler.GenerateSignature)

	// Start and run the server
	router.Run("localhost:8080")
}
