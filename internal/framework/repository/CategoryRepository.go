package repository

import (
	m "go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type categoryRepository struct {
	sqldb *gorm.DB
}

func NewCategoryRepository(sqldb *gorm.DB) *categoryRepository {
	return &categoryRepository{sqldb: sqldb}
}

func (br categoryRepository) FindAll() (chapt []response.Category, err error) {
	db := br.sqldb.Model(m.Category{}).Find(&chapt)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br categoryRepository) Create(chapt m.Category) (err error) {
	db := br.sqldb.Create(&chapt)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (br categoryRepository) Update(chapt m.Category) (err error) {
	db := br.sqldb.Updates(&chapt)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (br categoryRepository) Delete(id int) (err error) {
	db := br.sqldb.Delete(&m.Category{}, id)
	err = check.DBRecord(db, check.DELETE)
	return
}
