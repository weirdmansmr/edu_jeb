package main

import (
	"education/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
    var err error

    dsn := "host=localhost user=postgres dbname=education sslmode=disable password=Saykhanov01"
    DB, err = gorm.Open("postgres", dsn)
    if err != nil {
        panic(fmt.Sprintf("failed to connect database: %v", err))
    }

    DB.AutoMigrate(&models.Lesson{})
}
