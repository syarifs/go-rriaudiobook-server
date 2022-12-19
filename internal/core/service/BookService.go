package service

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/core/repository"
	"go-rriaudiobook-server/internal/utils"
	"go-rriaudiobook-server/internal/utils/errors"
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
	err = bs.repo.Create(m)
	return
}

func (bs BookService) Update(req request.BookRequest) (err error) {
	if req.ID == 0 {
		return errors.New(400, "No ID Provided")
	}
	m, _ := utils.TypeConverter[models.Book](req)
	m.ID = req.ID
	err = bs.repo.Update(m)
	return
}

func (bs BookService) Delete(id int) (err error) {
	if id == 0 {
		return errors.New(400, "No ID Provided")
	}
	err = bs.repo.Delete(id)
	return
}
