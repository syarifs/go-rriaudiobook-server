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

func (bs BookService) FindByUser(code string) (res []response.Book, err error) {
	res, err = bs.repo.FindByUser(code)
	return
}

func (bs BookService) Create(req request.BookRequest) (err error) {
	m, _ := utils.TypeConverter[models.Book](req)
	m.CoverImage, err = file.UploadFile(req.CoverImage, "cover/", "")
	if err != nil {
		return
	}

	err = bs.repo.Create(m)
	return
}

func (bs BookService) Update(req request.BookRequest) (err error) {
	if req.ID == 0 {
		return errors.New(400, "No ID Provided")
	}

	var cover string
	cover, err = bs.repo.GetCover(int(req.ID))
	if err == nil {
		err = file.RemoveFile(cover)
		if err != nil {
			return
		}
	}

	m, _ := utils.TypeConverter[models.Book](req)
	m.CoverImage, err = file.UploadFile(req.CoverImage, "cover/", "")
	if err != nil {
		return
	}

	err = bs.repo.Update(m)
	return
}

func (bs BookService) DeleteBook(id int) (err error) {
	if id == 0 {
		return errors.New(400, "No ID Provided")
	}

	var cover string
	cover, err = bs.repo.GetCover(id)
	if err == nil {
		err = file.RemoveFile(cover)
		if err != nil {
			return
		}
	}

	var chaptAudio []string
	chaptAudio, err = bs.repo.GetChapterAudioByBookID(id)
	if err == nil {
		for _, v := range chaptAudio {
			err = file.RemoveFile(v)
			if err != nil {
				return
			}
		}
	}

	err = bs.repo.Delete(id)
	return
}
