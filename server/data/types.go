package data

import (
    "gorm.io/gorm"
    "time"
)

// Model base struct for database models
type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Telemetry struct {
	Model
	CountryName string `json:"countryName"`
	CountryISO string `json:"countryIso"`
	TimeStamp time.Time `json:"timestamp"`
}

type TelemetryDto struct {
	CountryName string `json:"countryName"`
	CountryISO string `json:"countryIso"`
	TimeStamp time.Time `json:"timestamp"`
}

// ToTelemetry converts a TelemetryDto model to a Telemetry database model
func (dto TelemetryDto) ToTelemetry() *Telemetry {
	return &Telemetry{CountryName: dto.CountryName, CountryISO: dto.CountryISO, TimeStamp: dto.TimeStamp}
}

// Person is a database model
type Person struct {
	Model
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

// PersonDto is a dto model for the Person database model
type PersonDto struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

// ToPerson converts a PersonDto model to a Person database model
func (dto PersonDto) ToPerson() *Person {
	return &Person{FirstName: dto.FirstName, LastName: dto.LastName}
}
