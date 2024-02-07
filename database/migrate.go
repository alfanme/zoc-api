package database

import (
	"log"
	"zoc-api/models"
)

func AutoMigrate() {
	err := Client.AutoMigrate(&models.Customer{})
	
	if err != nil {
		log.Fatal("DB migration error", err.Error())
	}
}
