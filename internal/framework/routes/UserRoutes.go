package routes

import (
	"go-rriaudiobook-server/internal/framework/transport/controller"
	mw "go-rriaudiobook-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Group, acon *controller.UserController, middleware ...echo.MiddlewareFunc) {
	user := e.Group("/user", middleware...)
	user.POST("", acon.Create, mw.AdminPermission)
	user.GET("", acon.GetAllUser)
	user.GET("/:id", acon.GetUserByID)
	user.PUT("/:id/update", acon.Update, mw.AdminPermission)
	user.DELETE("/:id/delete", acon.Delete, mw.AdminPermission)
}
