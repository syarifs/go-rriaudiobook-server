package repository

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/request"
)

type CategoryRepository interface {
	FindAll() ([]request.CategoryRequest, error)
	Create(models.Category)
	Update(models.Category)
	Delete(int)
}
