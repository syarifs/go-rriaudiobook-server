package database

import (
	"fmt"
	"go-rriaudiobook-server/internal/utils/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPgSQL() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		config.DB_HOST, config.DB_USERNAME, config.DB_PASSWORD, config.DB_DATABASE, config.DB_PORT, config.DB_TIMEZONE,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}
