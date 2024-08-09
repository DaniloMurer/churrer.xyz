package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/data"
	"server/database"
)

// GetTelemetries handles the get request to retrieve all telemetries
func GetTelemetries(c *gin.Context) {
	telemetries := database.GetAllTelemetry()
	c.JSON(http.StatusOK, telemetries)
}

// CreateTelemetry handles a post request to create a new telemetry
func CreateTelemetry(c *gin.Context) {
	var newTelemetry data.TelemetryDto
	if err := c.BindJSON(&newTelemetry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		database.CreateTelemetry(newTelemetry.ToTelemetry())
		c.JSON(http.StatusCreated, newTelemetry)
	}
}

// GetExperiences handles the get request to retrieve all experiences
func GetExperiences(c *gin.Context) {
	experiences := database.GetAllExperience()
	c.JSON(http.StatusOK, experiences)
}

// CreateExperience handles a post request to create a new experience
func CreateExperience(c *gin.Context) {
	var newExperience data.ExperienceDto
	if err := c.BindJSON(&newExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		database.CreateExperience(newExperience.ToExperience())
		c.JSON(http.StatusCreated, newExperience)
	}
}
