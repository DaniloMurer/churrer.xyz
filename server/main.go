package main

import (
	"server/controller"
	"server/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dbErr := database.AutoMigration()
	if dbErr != nil {
		return
	}
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Accept", "Authorization", "Origin"}
	router.Use(cors.New(corsConfig))
	api := router.Group("/api")
	{
		api.GET("/telemetry", controller.GetTelemetries)
		api.POST("/telemetry", controller.CreateTelemetry)

		api.GET("/experience", controller.GetExperiences)
		api.POST("/experience", controller.CreateExperience)
		api.DELETE("/experience/:id", controller.DeleteExperience)
		api.PUT("/experience", controller.UpdateExperience)

		api.GET("/technology", controller.GetTechnologies)
		api.POST("/technology", controller.CreateTechnology)
		api.DELETE("/technology/:id", controller.DeleteTechnology)
		api.PUT("/technology", controller.UpdateTechnology)

		api.POST("/authentication", controller.AuthenticateUser)
	}
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		panic("we're fucked")
	}
}
