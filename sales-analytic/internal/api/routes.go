package api

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/refresh", RefreshDataHandler)
    r.GET("/analytics/total-revenue", TotalRevenueHandler)
}
