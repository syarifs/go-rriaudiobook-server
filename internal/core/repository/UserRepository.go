package repository

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
)

type UserRepository interface {
	FindAll() ([]response.User, error)
	FindByID(int) (response.UserDetail, error)
	Create(models.User) error
	Update(models.User) error
	Delete(int) error
}
