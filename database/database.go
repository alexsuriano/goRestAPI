package database

import (
	"log"
	"time"

	"github.com/alexsuriano/goRestAPI/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := db.DB()
	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

//Get a database connection started in the main.go
func GetDatabase() *gorm.DB {
	return db
}
