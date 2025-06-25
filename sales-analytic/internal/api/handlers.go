package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "sales-analytics/db"           // import your DB connection package
    "sales-analytics/internal/parser" // import your CSV loader
)


func RefreshDataHandler(c *gin.Context) {
    err := parser.LoadCSVToDB(db.DB, "data/sales.csv")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to refresh data", "error": err.Error()})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "Data refreshed successfully"})
    }
}

func TotalRevenueHandler(c *gin.Context) {
    var total float64
    start := c.Query("start")
    end := c.Query("end")

    result := db.DB.Raw(`
        SELECT SUM((unit_price * quantity_sold) * (1 - discount)) 
        FROM orders 
        WHERE date_of_sale BETWEEN ? AND ?`, start, end).Scan(&total)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to calculate revenue", "error": result.Error.Error()})
        return
    }
    
    
    c.JSON(200, gin.H{"total_revenue": total})
}
