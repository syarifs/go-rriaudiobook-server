package database

import (
	"go-rriaudiobook-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		models.User{},
		models.Book{},
		models.Category{},
		models.Chapter{},
	)
	return
}
