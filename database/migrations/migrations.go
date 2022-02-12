package migrations

import (
	"github.com/alexsuriano/goRestAPI/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
