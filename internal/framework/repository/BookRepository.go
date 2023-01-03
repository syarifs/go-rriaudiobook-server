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
		Scan(&res)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br bookRepository) FindByUser(code string) (res []response.Book, err error) {
	db := br.sqldb.Model(&m.Book{}).
		Select(`books.*`, `"Category".name as category`).
		Where("books.user_code = ?", code).
		Joins(`Category`).
		Scan(&res)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br bookRepository) FindByID(id int) (res response.BookDetail, err error) {
	db := br.sqldb.Model(&m.Book{}).
		Select(`books.*`, `"Category".name as category`).
		Joins(`Category`).
		Where(`books.id = ?`, id).
		Scan(&res)

	br.sqldb.Model(&m.Chapter{}).
		Where(`book_id = ?`, id).
		Scan(&res.Chapter)

	err = check.DBRecord(db, check.FIND)
	return
}

func (br bookRepository) Create(book m.Book) (err error) {
	db := br.sqldb.Create(&book)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (br bookRepository) Update(book m.Book) (err error) {
	db := br.sqldb.Updates(&book)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (br bookRepository) GetCover(id int) (file string, err error) {
	db := br.sqldb.Model(&m.Book{}).Where("id = ?", id).Select(`cover_image`).Scan(&file)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br bookRepository) GetChapterAudioByBookID(book_id int) (file []string, err error) {
	db := br.sqldb.Model(&m.Chapter{}).
		Where("book_id = ?", book_id).
		Select(`media_path`).Scan(&file)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br bookRepository) Delete(id int) (err error) {
	db := br.sqldb.Delete(&m.Book{}, id)
	err = check.DBRecord(db, check.DELETE)
	return
}
