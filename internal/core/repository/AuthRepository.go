package repository

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
)

type AuthRepository interface {
	FindByCode(string) (response.User, error)
	Register(models.User) error
	FindByEmail(string) (response.User, error)
	ChangePassword(models.User) error
	SaveToken(models.Token) error
	UpdateToken(models.Token, models.Token) error
	RevokeToken(string) error
}
