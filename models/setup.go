package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbURL := "postgres://postgres:mysecretpassword@localhost:5432/go_rest_api"

	database, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
