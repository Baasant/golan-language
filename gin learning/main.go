package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new router with default middleware
	router := gin.Default()

	// Define a route handler for GET requests to "/hello"
	router.GET("/hello", func(c *gin.Context) {
		// Set the response status code to 200 (OK)
		c.Status(http.StatusOK)

		// Return a JSON response with a "message" field
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello bassant",
		})
	})

	// Start the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
