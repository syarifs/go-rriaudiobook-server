package service

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/core/repository"
	"go-rriaudiobook-server/internal/utils"
	"go-rriaudiobook-server/internal/utils/errors"
	"go-rriaudiobook-server/internal/utils/file"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (bs BookService) FindAll() (res []response.Book, err error) {
	res, err = bs.repo.FindAll()
	return
}

func (bs BookService) FindByID(id int) (res response.BookDetail, err error) {
	res, err = bs.repo.FindByID(id)
	return
}

func (bs BookService) Create(req request.BookRequest) (err error) {
	m, _ := utils.TypeConverter[models.Book](req)
	err = bs.repo.CreateBook(m)
	return
}

func (bs BookService) Update(id int, req request.BookRequest) (err error) {
	if id == 0 {
		return errors.New(400, "No ID Provided")
	}

	var cover string
	cover, err = bs.repo.GetBookCover(id)
	if err != nil {
		return
	}

	if cover != "" {
		err = file.RemoveFile(cover)
		if err != nil {
			return
		}
	}

	m, _ := utils.TypeConverter[models.Book](req)
	m.ID = uint(id)
	err = bs.repo.UpdateBook(m)
	return
}

func (bs BookService) InsertChapter(book_id int, req request.ChapterRequest) (err error) {
	if book_id == 0 {
		return errors.New(400, "No ID Provided")
	}

	m, _ := utils.TypeConverter[models.Chapter](req)
	m.BookID = uint(book_id)
	err = bs.repo.InsertChapter(m)
	return
}

func (bs BookService) UpdateChapter(book_id, chapter_id int, req request.ChapterRequest) (err error) {
	if book_id == 0 {
		return errors.New(400, "No ID Provided")
	}

	var media_file string
	media_file, err = bs.repo.GetChapterAudio(book_id, chapter_id)
	if err != nil {
		return
	}

	if media_file != "" {
		err = file.RemoveFile(media_file)
		if err != nil {
			return
		}
	}

	m, _ := utils.TypeConverter[models.Chapter](req)
	m.BookID = uint(book_id)
	m.Code = uint(chapter_id)
	err = bs.repo.UpdateChapter(m)
	return
}

func (bs BookService) DeleteBook(id int) (err error) {
	if id == 0 {
		return errors.New(400, "No ID Provided")
	}

	var cover string
	cover, err = bs.repo.GetBookCover(id)
	if err != nil {
		return
	}

	if cover != "" {
		err = file.RemoveFile(cover)
		if err != nil {
			return
		}
	}

	var chaptAudio []string
	chaptAudio, err = bs.repo.GetChapterAudioByBookID(id)
	if err != nil {
		return
	}

	for _, v := range chaptAudio {
		err = file.RemoveFile(v)
		if err != nil {
			return
		}
	}

	err = bs.repo.DeleteBook(id)
	return
}

func (bs BookService) DeleteChapter(book_id, chapter_id int) (err error) {
	if book_id == 0 {
		return errors.New(400, "No ID Provided")
	}

	var cover string
	cover, err = bs.repo.GetChapterAudio(book_id, chapter_id)
	if err != nil {
		return
	}

	err = file.RemoveFile(cover)
	if err != nil {
		return
	}

	err = bs.repo.DeleteChapter(book_id, chapter_id)
	return
}
