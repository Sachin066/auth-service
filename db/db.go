package db

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

)


var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=user password=pass dbname=auth_db port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }
}
