package main

import (
	"go-rriaudiobook-server/internal/core/service"
	"go-rriaudiobook-server/internal/framework/database"
	"go-rriaudiobook-server/internal/framework/repository"
	"go-rriaudiobook-server/internal/framework/routes"
	"go-rriaudiobook-server/internal/framework/transport/controller"
	mw "go-rriaudiobook-server/internal/framework/transport/middleware"
	"go-rriaudiobook-server/internal/utils/config"
	"go-rriaudiobook-server/internal/utils/jwt"
	"go-rriaudiobook-server/internal/utils/logger"
	"go-rriaudiobook-server/internal/utils/validators"

	_ "go-rriaudiobook-server/docs"

	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           RRI Audiobook API
// @version         1.0
// @description     server API for RRI Audiobook Application.

// @securityDefinitions.apikey ApiKey
// @in header
// @name Authorization

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host     go-rriaudiobook-server-production.up.railway.app
// @BasePath  /api
// @schemes http https
func main() {

	config.LoadConfig()

	db, mongodb := database.InitDatabase()
	repo := repository.NewRepository(db, mongodb)
	serv := service.NewService(repo)
	ctrl := controller.NewController(serv)
	logger.NewLogger(mongodb)
	validators.NewValidator(db)

	e := echo.New()
	e.Use(emw.CORS())
	e.GET("/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	jwt.NewJWTConnection(mongodb)
	routes.NewRoutes(api, ctrl, mw.JWT)

	mw.Logging(e)

	e.Start(":" + config.SERVER_PORT)
}
