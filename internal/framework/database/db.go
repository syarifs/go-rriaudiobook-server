package database

import (
	"go-rriaudiobook-server/internal/framework/database/seeds"
	"go-rriaudiobook-server/internal/utils/config"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func InitDatabase() (sqldb *gorm.DB, mongodb *mongo.Database) {
	var err error

	mongodb = initMongoDB()

	if config.DB_DRIVER == "mysql" {
		sqldb, err = initMySQL()
		if err != nil {
			panic(err)
		}
	} else if config.DB_DRIVER == "pgsql" {
		sqldb, err = initPgSQL()
		if err != nil {
			panic(err)
		}
	} else {
		panic("SQL Database Not Supported")
	}

	err = migrateDB(sqldb)
	if err != nil {
		panic(err)
	}

	seeds.NewSeeders(sqldb)

	return
}
