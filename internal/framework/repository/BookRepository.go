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

func (br bookRepository) CreateBook(book m.Book) (err error) {
	db := br.sqldb.Create(&book)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (br bookRepository) InsertChapter(chapt m.Chapter) (err error) {
	db := br.sqldb.Create(&chapt)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (br bookRepository) UpdateBook(book m.Book) (err error) {
	db := br.sqldb.Updates(&book)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (br bookRepository) GetBookCover(id int) (file string, err error) {
	db := br.sqldb.Model(&m.Book{}).Where("id = ?", id).Select(`cover_image`).Scan(&file)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br bookRepository) GetChapterAudio(book_id, chapter_id int) (file string, err error) {
	db := br.sqldb.Model(&m.Chapter{}).
		Where("book_id = ?", book_id).
		Where("code = ?", chapter_id).
		Select(`media_path`).Scan(&file)
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

func (br bookRepository) DeleteBook(id int) (err error) {
	db := br.sqldb.Delete(&m.Book{}, id)
	err = check.DBRecord(db, check.DELETE)
	return
}

func (br bookRepository) UpdateChapter(chapt m.Chapter) (err error) {
	db := br.sqldb.Model(&m.Chapter{}).
		Where("book_id = ?", chapt.BookID).Where("code = ?", chapt.Code).
		Updates(&chapt)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (br bookRepository) DeleteChapter(chapter_id, book_id int) (err error) {
	db := br.sqldb.Delete(&m.Chapter{}).
		Where(`code = ?`, chapter_id).
		Where(`book_id = ?`, chapter_id)
	err = check.DBRecord(db, check.DELETE)
	return
}
