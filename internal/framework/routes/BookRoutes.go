package routes

import (
	"go-rriaudiobook-server/internal/framework/transport/controller"
	// mw "go-rriaudiobook-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewBookRoutes(e *echo.Group, acon *controller.BookController, middleware ...echo.MiddlewareFunc) {
	patient := e.Group("/books", middleware...)
	patient.POST("", acon.Create)
	patient.GET("", acon.GetAllBook)
	patient.GET("/:id", acon.GetBookByID)
	// patient.PUT("/:id/update", acon.Update, mw.AdminPermission)
	// patient.DELETE("/:id/delete", acon.Delete, mw.AdminPermission)
}
