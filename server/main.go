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
	}
	err := router.Run("localhost:8080")
	if err != nil {
		panic("we're fucked")
	}
}
