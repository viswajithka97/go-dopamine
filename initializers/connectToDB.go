package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var db *gorm.DB

func ConnectDB() {

	var err error
	dsn := os.Getenv("DB_URL")

	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB")

	} else {
		fmt.Println("Connected to DB")
	}
}
