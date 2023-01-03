package database

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/utils/config"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	if config.TOKEN_DRIVER == "sql" {
		db.AutoMigrate(models.Token{})
	}
	err = db.AutoMigrate(
		models.User{},
		models.Category{},
		models.Book{},
		models.Chapter{},
	)
	return
}
