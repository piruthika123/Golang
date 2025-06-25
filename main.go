package main

import (
    "sales-analytics/db"
    "sales-analytics/internal/api"
    "sales-analytics/internal/models"
    "sales-analytics/internal/refresh"

    "github.com/gin-gonic/gin"
)

func main() {
    // Connect to the database
    dbConn := db.ConnectDB()

    // Run auto migration
    dbConn.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{})

    // Start background refresh job
    refresh.StartDailyRefreshJob()

    // Initialize and register routes
    r := gin.Default()
    api.RegisterRoutes(r)

    // Run the server
    r.Run(":8080")
}
