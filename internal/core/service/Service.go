package service

import "go-rriaudiobook-server/internal/core/repository"

type Service struct {
	Auth *AuthService
	User *UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth),
		User: NewUserService(repo.User),
	}
}
