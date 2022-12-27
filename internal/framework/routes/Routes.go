package routes

import (
	"go-rriaudiobook-server/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewRoutes(e *echo.Group, ctrl *controller.Controller, middleware ...echo.MiddlewareFunc) {
	NewAuthRoutes(e, ctrl.Auth)
	NewUserRoutes(e, ctrl.User, middleware...)
	NewBookRoutes(e, ctrl.Book, middleware...)
	NewChapterRoutes(e, ctrl.Chapter, middleware...)
	NewCategoryRoutes(e, ctrl.Category, middleware...)
}
