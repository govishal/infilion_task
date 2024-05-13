package main

import (
	"project/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define endpoint handler for /person/:person_id/info
	router.GET("/person/:person_id/info", handler.GetPersonInfoHandler)
	router.POST("/person/create", handler.AddPersonInfoHandler)

	// Run the server on port 8080
	router.Run(":8080")

	// Run the server on port 8080
	router.Run(":8080")
}