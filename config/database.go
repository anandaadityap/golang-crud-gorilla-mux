package config

import (
	"fmt"
	"golang-crud/models"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connStr := fmt.Sprintf("%v", ENV.DB_CONNECTION)

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.Books{}, &models.Author{})

	DB = db
	log.Println("Database connected")
}
