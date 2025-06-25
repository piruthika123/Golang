package db

import (
    "fmt"
    "log"
    "sales-analytics/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect initializes the DB connection using GORM and the config
func Connect() {
    cfg := config.AppConfig

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Warn),
    })
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("✅ Successfully connected to database")
}
