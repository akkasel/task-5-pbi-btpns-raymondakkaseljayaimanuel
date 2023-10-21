package database

import (
	"log"
	"task5-pbi-btpns-raymondakkaseljayaimanuel/models"
)

func Migrate() {
	Instance.AutoMigrate(&models.User{}, &models.Photo{})
	log.Println("Database Migration Completed!")
}
