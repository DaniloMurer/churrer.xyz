package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
    "server/data"
)

var connectionString = "host=localhost user=postgres password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Zurich"
var database *gorm.DB

var logger = log.New(os.Stdout, "[XYZ] - ", log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

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
	return database.AutoMigrate(&data.Person{}, &data.Telemetry{})
}

// GetAllPerson returns all people from the database
func GetAllPerson() []data.Person {
	logger.Println("Retrieving all people")
    var people []data.Person
	database.Find(&people)
	return people
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

// CreatePerson creates new data.Person entry in database
func CreatePerson(person *data.Person) {
	logger.Printf("Creating Person: {%v %v}", person.FirstName, person.LastName)
	database.Create(person)
}