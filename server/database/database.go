package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"server/data"
)

var connectionString = os.Getenv("DB_CONNECTION")
var database *gorm.DB

var logger = log.New(os.Stdout, "[XYZ] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

// openDatabaseConnection creates a postgresql connection
func openDatabaseConnection() {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		logger.Fatal("Error while opening database connection", err.Error())
	} else {
		database = db
	}
}

// AutoMigration migrates the schema defined in the data package to the database
func AutoMigration() error {
	openDatabaseConnection()
	return database.AutoMigrate(&data.Telemetry{}, &data.Experience{}, &data.Technology{})
}

// GetAllTelemetry returns all telemetries from the database
func GetAllTelemetry() []data.Telemetry {
	logger.Println("Retrieving all telemetries")
	var telemetries []data.Telemetry
	database.Find(&telemetries)
	return telemetries
}

// CreateTelemetry creates new data.Telemetry entry in database
func CreateTelemetry(telemetry *data.Telemetry) {
	logger.Println("Creating new telemetry entry")
	database.Create(telemetry)
}

// GetAllExperience returns all experiences from the database
func GetAllExperience() []data.Experience {
	logger.Println("Retrieving all experiences")
	var experiences []data.Experience
	database.Find(&experiences)
	return experiences
}

// CreateExperience creates new data.Experience entry in database
func CreateExperience(experience *data.Experience) {
	logger.Println("Creating new experience entry")
	database.Create(experience)
}

// GetAllTechnology return all technologies from the database
func GetAllTechnology() []data.Technology {
	logger.Println("Retrieving all technologies")
	var technologies []data.Technology
	database.Find(&technologies)
	return technologies
}

// CreateTechnology creates new data.Technology entry in database
func CreateTechnology(technology *data.Technology) {
	logger.Println("Creating new technology entry")
	database.Create(technology)
}
