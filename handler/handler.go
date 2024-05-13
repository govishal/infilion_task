package handler

import (
	"net/http"
	"project/logic"
	"project/model"

	"github.com/gin-gonic/gin"
)

func GetPersonInfoHandler(c *gin.Context) {
	// Get person ID from request URL
	personID := c.Param("person_id")

	// Fetch person info based on person ID using logic function
	person, err := logic.GetPerson(personID)
	if err != nil {
		// If an error occurs (e.g., person not found), return an error response
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	// If person info is fetched successfully, return it in the response
	c.JSON(http.StatusOK, person)
}


func AddPersonInfoHandler(c *gin.Context) {
	// Get person ID from request URL
	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Fetch person info based on person ID using logic function
	_, err := logic.AddPerson(person)
	if err != nil {
		// If an error occurs (e.g., person not found), return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add person"})
		return
	}

	// If person info is fetched successfully, return it in the response
	c.JSON(http.StatusOK, person)
}