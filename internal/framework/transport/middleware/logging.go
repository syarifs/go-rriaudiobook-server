package middleware

import (
	"go-rriaudiobook-server/internal/utils/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Logging(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
		Output: &logger.LogDriver,
	}))
	e.Logger.SetLevel(log.DEBUG)
}
