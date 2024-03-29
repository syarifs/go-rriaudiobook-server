package routes

import (
	"go-rriaudiobook-server/internal/framework/transport/controller"
	mw "go-rriaudiobook-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewChapterRoutes(e *echo.Group, acon *controller.ChapterController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.UploaderPermission)
	book := e.Group("/books")
	book.POST("/:book_id/chapters", acon.InsertChapter, middleware...)
	book.PUT("/:book_id/chapters/:chapter_id/update", acon.UpdateChapter, middleware...)
	book.DELETE("/:book_id/chapters/:chapter_id/delete", acon.DeleteChapter, middleware...)
}
