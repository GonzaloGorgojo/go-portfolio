package database

import (
	"log"

	model "github.com/GonzaloGorgojo/go-backend-portfolio/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(model.Experience{})

	return db
}
