package database

import (
	"fmt"
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/utils/config"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	fmt.Println(config.TOKEN_DRIVER)
	if config.TOKEN_DRIVER == "sql" {
		db.AutoMigrate(models.Token{})
	}
	err = db.AutoMigrate(
		models.User{},
		models.Book{},
		models.Category{},
		models.Chapter{},
	)
	return
}
