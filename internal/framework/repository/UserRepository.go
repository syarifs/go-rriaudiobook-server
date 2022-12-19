package repository

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type userRepository struct {
	sqldb *gorm.DB
}

func NewUserRepository(sqldb *gorm.DB) *userRepository {
	return &userRepository{sqldb: sqldb}
}

func (repo *userRepository) FindAll() (res []response.User, err error) {
	db := repo.sqldb.Model(&models.User{}).
		Select(`users.*`, `"Role".name as role`).
		Joins("Role").
		Find(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo *userRepository) FindByID(id int) (res response.UserDetail, err error) {
	db := repo.sqldb.Model(&models.User{}).
		Select(`users.*, "Role".name as role`).
		Joins("Role").
		Where("users.id = ?", id).
		Scan(&res)

	err = check.DBRecord(db, check.FIND)
	return
}

func (repo *userRepository) Create(us models.User) (err error) {
	db := repo.sqldb.Create(&us)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (repo *userRepository) Update(us models.User) (err error) {
	db := repo.sqldb.Updates(&us)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (repo *userRepository) Delete(id int) (err error) {
	db := repo.sqldb.Delete(&models.User{}, id)
	err = check.DBRecord(db, check.DELETE)
	return
}
