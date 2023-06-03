package main

import (
	"log"
	"yaonkey/inhouse/pkg/database"
	"yaonkey/inhouse/pkg/models"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	db := database.GetDB()
	db.Create(&models.Admin{})
	db.Create(&models.Site{})
}
