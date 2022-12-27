package repository

import (
	m "go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/utils/errors/check"

	"gorm.io/gorm"
)

type chapterRepository struct {
	sqldb *gorm.DB
}

func NewChapterRepository(sqldb *gorm.DB) *chapterRepository {
	return &chapterRepository{sqldb: sqldb}
}

func (br chapterRepository) Create(chapt m.Chapter) (err error) {
	db := br.sqldb.Create(&chapt)
	err = check.DBRecord(db, check.CREATE)
	return
}

func (br chapterRepository) GetAudio(chapter_id, book_id int) (file string, err error) {
	db := br.sqldb.Model(&m.Chapter{}).
		Where("book_id = ?", book_id).
		Where("code = ?", chapter_id).
		Select(`media_path`).Scan(&file)
	err = check.DBRecord(db, check.FIND)
	return
}

func (br chapterRepository) Update(chapt m.Chapter) (err error) {
	db := br.sqldb.Model(&m.Chapter{}).
		Where("book_id = ?", chapt.BookID).Where("code = ?", chapt.Code).
		Updates(&chapt)
	err = check.DBRecord(db, check.UPDATE)
	return
}

func (br chapterRepository) Delete(chapter_id, book_id int) (err error) {
	db := br.sqldb.Delete(&m.Chapter{}, "code = ? AND book_id = ?", chapter_id, book_id)
	err = check.DBRecord(db, check.DELETE)
	return
}
