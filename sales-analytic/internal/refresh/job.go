package job

import (
	"log"
	"os"
	"time"
	"sales-analytics/db"               //  import the DB package
    "sales-analytics/internal/parser"  //  assuming LoadCSVToDB is in parser

)

func StartDailyRefreshJob() {
    ticker := time.NewTicker(24 * time.Hour)
    go func() {
        for {
            <-ticker.C
            err := parser.LoadCSVToDB(db.DB, "data/sales.csv")
            if err != nil {
                logToFile("Failed refresh: " + err.Error())
            } else {
                logToFile("Successful refresh")
            }
        }
    }()
}

func logToFile(message string) {
    f, _ := os.OpenFile("logs/refresh.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    logger := log.New(f, "", log.LstdFlags)
    logger.Println(message)
}
