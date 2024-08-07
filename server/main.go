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
    router.GET("/api/telemetry", controller.GetTelemetries)
    router.POST("/api/telemetry", controller.CreateTelemetry)
    err := router.Run("localhost:8080")
    if err != nil {
        panic("we're fucked")
    }
}