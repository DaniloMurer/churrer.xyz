package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "server/data"
    "server/database"
)

// GetPeople handles the get request to retrieve all people
func GetPeople(c *gin.Context) {
    people := database.GetAllPerson()
    c.JSON(http.StatusOK, people)
}

// CreatePerson handles a post request to create a new person
func CreatePerson(c *gin.Context) {
    var newPerson data.PersonDto
    if err := c.BindJSON(&newPerson); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err})
    } else {
        database.CreatePerson(newPerson.ToPerson())
        c.JSON(http.StatusCreated, gin.H{})
    }
}

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