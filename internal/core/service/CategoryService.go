package service

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/core/repository"
	"go-rriaudiobook-server/internal/utils"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (cs CategoryService) FindAll() (res []response.Category, err error) {
	res, err = cs.repo.FindAll()
	return
}

func (cs CategoryService) Create(cat request.CategoryRequest) (err error) {
	var category models.Category
	category, _ = utils.TypeConverter[models.Category](cat)
	err = cs.repo.Create(category)
	return
}

func (cs CategoryService) Update(cat request.CategoryRequest) (err error) {
	var category models.Category
	category, _ = utils.TypeConverter[models.Category](cat)
	err = cs.repo.Update(category)
	return
}

func (cs CategoryService) Delete(id int) (err error) {
	err = cs.repo.Delete(id)
	return
}
