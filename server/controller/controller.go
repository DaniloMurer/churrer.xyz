package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"server/data"
	"server/database"
	"strings"
)

// Authorize validates if the provided header contains valid credentials for authorization.
// It expects the header to be in the format "Basic base64encodedCredentials".
// If the header does not follow the expected format, or if the credentials are incorrect,
// it returns false. Otherwise, it returns true.
func Authorize(header string) bool {
	basicAuthHeader := strings.TrimPrefix(header, "Basic ")
	decodedHeader, err := base64.StdEncoding.DecodeString(basicAuthHeader)
	if err != nil {
		return false
	}
	credentials := strings.Split(string(decodedHeader), ":")
	if len(credentials) != 2 {
		return false
	}
	username := credentials[0]
	password := credentials[1]
	if username != os.Getenv("ADMIN_USERNAME") || password != os.Getenv("ADMIN_PASSWORD") {
		return false
	}
	return true
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

// GetExperiences handles the get request to retrieve all experiences
func GetExperiences(c *gin.Context) {
	experiences := database.GetAllExperience()
	c.JSON(http.StatusOK, experiences)
}

// CreateExperience handles a post request to create a new experience
func CreateExperience(c *gin.Context) {
	encodedHeader := c.GetHeader("Authorization")
	authorized := Authorize(encodedHeader)
	if !authorized {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	var newExperience data.ExperienceDto
	if err := c.BindJSON(&newExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		database.CreateExperience(newExperience.ToExperience())
		c.JSON(http.StatusCreated, newExperience)
	}
}
