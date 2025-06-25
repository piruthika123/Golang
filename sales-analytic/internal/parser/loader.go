package loader

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sales-analytic/models"
)

func LoadCSVToDB(db *gorm.DB, filepath string) error {
    f, err := os.Open(filepath)
    if err != nil {
        return err
    }
    defer f.Close()

    reader := csv.NewReader(f)
    reader.Read() // skip header

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }

        customer := models.Customer{CustomerID: record[2], Name: record[12], Email: record[13], Address: record[14]}
        product := models.Product{ProductID: record[1], Name: record[3], Category: record[4]}

        date, _ := time.Parse("2006-01-02", record[6])
        quantity, _ := strconv.Atoi(record[7])
        unitPrice, _ := strconv.ParseFloat(record[8], 64)
        discount, _ := strconv.ParseFloat(record[9], 64)
        shipping, _ := strconv.ParseFloat(record[10], 64)

        order := models.Order{
            OrderID: record[0], ProductID: product.ProductID, CustomerID: customer.CustomerID,
            Region: record[5], DateOfSale: date, QuantitySold: quantity,
            UnitPrice: unitPrice, Discount: discount, ShippingCost: shipping,
            PaymentMethod: record[11],
        }

        db.Clauses(clause.OnConflict{DoNothing: true}).Create(&customer)
        db.Clauses(clause.OnConflict{DoNothing: true}).Create(&product)
        db.Clauses(clause.OnConflict{DoNothing: true}).Create(&order)
    }

    return nil
}
