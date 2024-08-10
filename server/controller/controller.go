package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"server/data"
	"server/database"
)

var adminUsername = os.Getenv("ADMIN_USERNAME")
var adminPassword = os.Getenv("ADMIN_PASSWORD")

var logger = log.New(os.Stdout, "[XYZ] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

// Authorize checks if the provided username and password match the admin credentials.
// Returns true if authorized, otherwise false. If there is an error, returns an error object.
func Authorize(username string, password string, ok bool) (bool, error) {
	if !ok {
		logger.Println("WARN: Non compliant BasicAuth header received")
		return false, fmt.Errorf("non compliant BasicAuth Header found")
	}
	if username != adminUsername || password != adminPassword {
		return false, nil
	}
	return true, nil
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
		logger.Printf("ERROR: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	database.CreateTelemetry(newTelemetry.ToTelemetry())
	c.JSON(http.StatusCreated, newTelemetry)
}

// GetExperiences handles the get request to retrieve all experiences
func GetExperiences(c *gin.Context) {
	experiences := database.GetAllExperience()
	c.JSON(http.StatusOK, experiences)
}

// CreateExperience handles a post request to create a new experience
func CreateExperience(c *gin.Context) {
	authorized, err := Authorize(c.Request.BasicAuth())
	if err != nil {
		logger.Printf("ERROR: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if !authorized {
		logger.Println("WARN: Attempted login with wrong credentials")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	var newExperience data.ExperienceDto
	if err := c.BindJSON(&newExperience); err != nil {
		logger.Printf("ERROR: %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	database.CreateExperience(newExperience.ToExperience())
	c.JSON(http.StatusCreated, newExperience)
}

// GetTechnologies handles the get request to retrieve all technologies
func GetTechnologies(c *gin.Context) {
	technologies := database.GetAllTechnology()
	c.JSON(http.StatusOK, technologies)
}

// CreateTechnology handles the post request to create a new technology
func CreateTechnology(c *gin.Context) {
	authorized, err := Authorize(c.Request.BasicAuth())
	if err != nil {
		logger.Printf("ERROR: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if !authorized {
		logger.Println("WARN: Attempted login with wrong credentials")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	var newTechnology data.TechnologyDto
	if err := c.BindJSON(&newTechnology); err != nil {
		logger.Printf("ERROR: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	database.CreateTechnology(newTechnology.ToTechnology())
	c.JSON(http.StatusCreated, newTechnology)
}
