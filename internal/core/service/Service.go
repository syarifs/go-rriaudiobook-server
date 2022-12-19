package service

import "go-rriaudiobook-server/internal/core/repository"

type Service struct {
	Auth *AuthService
	User *UserService
	Book *BookService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth),
		User: NewUserService(repo.User),
		Book: NewBookService(repo.Book),
	}
}
