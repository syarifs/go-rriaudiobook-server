package routes

import (
	"go-rriaudiobook-server/internal/framework/transport/controller"
	mw "go-rriaudiobook-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewCategoryRoutes(e *echo.Group, acon *controller.CategoryController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.UploaderPermission)

	book := e.Group("/categories")
	book.POST("", acon.Create, middleware...)
	book.GET("", acon.FindAll)
	book.PUT("/:id/update", acon.Update, middleware...)
	book.DELETE("/:id/delete", acon.Delete, middleware...)
}
