package database

import (
	"log"
	"zoc-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	
	if err != nil {
		log.Fatal("DB connection error", err.Error())
	}

	Client = db
}