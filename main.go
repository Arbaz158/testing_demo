package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// Access the environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	apiKey := os.Getenv("API_KEY")

	fmt.Println("Database URL:", databaseURL)
	fmt.Println("API Key:", apiKey)

	r := gin.Default()
	r.GET("/get", HelloWorld)
	r.Run(":8080")
}

func HelloWorld(c *gin.Context) {
	fmt.Println("Hello World")
	c.JSON(http.StatusOK, gin.H{"message": "Hii this is testing of docker deployment\n"})
	c.JSON(http.StatusOK, gin.H{"message": "just checking for some enhancement\n"})
	c.JSON(http.StatusOK, gin.H{"message": "THird time modifying something"})
}
