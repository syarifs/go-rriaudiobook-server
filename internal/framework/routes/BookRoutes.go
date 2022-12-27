package routes

import (
	"go-rriaudiobook-server/internal/framework/transport/controller"
	mw "go-rriaudiobook-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewBookRoutes(e *echo.Group, acon *controller.BookController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.UploaderPermission)

	book := e.Group("/books")
	book.POST("", acon.InsertBook, middleware...)
	book.GET("", acon.GetAllBook)
	book.GET("/:id", acon.GetBookByID)
	book.PUT("/:id/update", acon.UpdateBook, middleware...)
	book.DELETE("/:id/delete", acon.DeleteBook, middleware...)
}
