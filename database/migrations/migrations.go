package migrations

import (
	"github.com/hjunor/api-rest-golang.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
