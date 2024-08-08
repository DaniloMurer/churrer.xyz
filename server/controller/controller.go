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
