package initializers

import (
	"inspection-api/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.ImageField{})
}