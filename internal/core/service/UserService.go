package service

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/core/repository"
	"go-rriaudiobook-server/internal/utils"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (srv UserService) Create(req request.UserRequest) (err error) {
	m, _ := utils.TypeConverter[models.User](req)
	m.Password, _ = utils.HashPassword(req.Password)
	m.Status = 1
	err = srv.repo.Create(m)
	return
}

func (srv UserService) Update(req request.UserRequest) (err error) {
	m, _ := utils.TypeConverter[models.User](req)
	m.Password, _ = utils.HashPassword(req.Password)
	m.ID = uint(req.ID)
	err = srv.repo.Update(m)
	return
}

func (srv UserService) Delete(id int) (err error) {
	err = srv.repo.Delete(id)
	return
}

func (srv UserService) FindAll() (user []response.User, err error) {
	user, err = srv.repo.FindAll()
	return
}

func (srv UserService) FindByID(id int) (user response.UserDetail, err error) {
	user, err = srv.repo.FindByID(id)
	return
}
