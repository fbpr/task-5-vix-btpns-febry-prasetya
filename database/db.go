package database

import (
	"fmt"
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "developer"
	password = "supersecretpassword"
	dbPort   = "5432"
	dbName   = "vixbtpns"
)

var (
	db *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}

func GetDb() *gorm.DB {
	return db
}