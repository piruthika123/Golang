package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config holds the global application configuration
type Config struct {
    DBUser     string
    DBPassword string
    DBName     string
    DBHost     string
    DBPort     string
    ServerPort string
    LogPath    string
}

// AppConfig stores the loaded configuration
var AppConfig Config

// LoadConfig loads environment variables and initializes AppConfig
func LoadConfig() {
    // Load from .env file if available
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, loading from system environment")
    }

    AppConfig = Config{
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "sales_db"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        ServerPort: getEnv("SERVER_PORT", "8080"),
        LogPath:    getEnv("LOG_PATH", "./logs/refresh.log"),
    }
}

// getEnv fetches environment variable or returns default
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
