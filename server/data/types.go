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