package routes

import (
	"go-rriaudiobook-server/internal/framework/transport/controller"
	// mw "go-rriaudiobook-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewBookRoutes(e *echo.Group, acon *controller.BookController, middleware ...echo.MiddlewareFunc) {
	patient := e.Group("/books", middleware...)
	patient.POST("", acon.InsertBook)
	patient.GET("", acon.GetAllBook)
	patient.GET("/:id", acon.GetBookByID)
	patient.PUT("/:id/update", acon.UpdateBook)
	patient.DELETE("/:id/delete", acon.DeleteBook)
	patient.POST("/:book_id/chapters", acon.InsertChapter)
	patient.PUT("/:book_id/chapters/:chapter_id/update", acon.UpdateChapter)
	patient.DELETE("/:book_id/chapters/:chapter_id/delete", acon.DeleteChapter)
}
