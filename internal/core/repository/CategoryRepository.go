package repository

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
)

type CategoryRepository interface {
	FindAll() ([]response.Category, error)
	Create(models.Category) error
	Update(models.Category) error
	Delete(int) error
}
