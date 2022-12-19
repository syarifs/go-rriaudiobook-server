package repository

import (
	m "go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type bookRepository struct {
	sqldb *gorm.DB
}

func NewBookRepository(sqldb *gorm.DB) *bookRepository {
	return &bookRepository{sqldb: sqldb}
}

func (br bookRepository) FindAll() (res []response.Book, err error) {
	db := br.sqldb.Model(&m.Book{}).
		Select(`books.*`, `"Category".name as category`).
		Joins(`Category`).
		Find(&res)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br bookRepository) FindByID(id int) (res response.BookDetail, err error) {
	db := br.sqldb.Model(&m.Book{}).
		Select(`books.*`, `"Category".name as category`).
		Joins(`Category`).
		Preload(`Chapter`).
		Where(`books.id = ?`, id).
		Scan(&res)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br bookRepository) Create(book m.Book) (err error) {
	db := br.sqldb.Create(&book)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (br bookRepository) InsertChapter(m.Chapter) (err error) {
	return
}

func (br bookRepository) Update(m.Book) (err error) {
	return
}

func (br bookRepository) Delete(int) (err error) {
	return
}
