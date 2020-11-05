package db

import (
	m "github.com/azdonald/sellerbase/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB global for accessing the database
var DB *gorm.DB

// InitDB sets up connection to the database
func InitDB() {
	dsn := "root:presario@tcp(127.0.0.1:3306)/sellerbase?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&m.Merchant{})
	DB = database
}
