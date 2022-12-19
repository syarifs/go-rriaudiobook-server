package controller

import "go-rriaudiobook-server/internal/core/service"

type Controller struct {
	Auth *AuthController
	User *UserController
	Book *BookController
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(srv.Auth),
		User: NewUserController(srv.User),
		Book: NewBookController(srv.Book),
	}
}
