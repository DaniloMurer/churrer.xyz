package main

import (
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/database"
)

func main() {
	dbErr := database.AutoMigration()
	if dbErr != nil {
		return
	}
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/telemetry", controller.GetTelemetries)
		api.POST("/telemetry", controller.CreateTelemetry)

		api.GET("/experience", controller.GetExperiences)
		api.POST("/experience", controller.CreateExperience)

		api.GET("/technology", controller.GetTechnologies)
		api.POST("/technology", controller.CreateTechnology)

		api.POST("/authentication", controller.AuthenticateUser)
	}
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("we're fucked")
	}
}
