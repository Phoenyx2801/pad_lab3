package db

import (
	"REST-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectPostgres() {

	dsn := os.Getenv("DB_CONNECT_STRING")

	//host=172.18.54.12 user=postgres password=postgres port=5432 sslmode=disable
	//"host=localhost user=postgres password=123qweASD port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.Book{})

	if err != nil {
		return
	}

	DB = db
}
