package middleware

import (
	"fmt"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/utils/config"
	"go-rriaudiobook-server/internal/utils/errors/check"
	"go-rriaudiobook-server/internal/utils/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwt.GetToken(c, jwt.ACCESS)

		fmt.Println("Token Get:", token)
		fmt.Println("Token Error:", err)
		fmt.Println("Access Time:", config.TOKEN_ACCESS_EXPIRE_TIME)
		fmt.Println("Refresh Time:", config.TOKEN_REFRESH_EXPIRE_TIME)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		err = jwt.FindToken(token)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		return next(c)
	}
}

func AdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := jwt.GetToken(c, jwt.ACCESS)
		role, err := jwt.GetTokenData(token, "role", jwt.ACCESS)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		if role != "administrator" && role != "admin" {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for administrator",
			})
		}

		return next(c)
	}
}

func UploaderPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := jwt.GetToken(c, jwt.ACCESS)
		role, err := jwt.GetTokenData(token, "role", jwt.ACCESS)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		if strings.ToLower(role.(string)) != "uploader" {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for uploader",
			})
		}

		return next(c)
	}
}

func AdminUploaderPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := jwt.GetToken(c, jwt.ACCESS)
		role, err := jwt.GetTokenData(token, "role", jwt.ACCESS)

		if r, ok := check.HTTP(nil, err, "Validate Token"); !ok {
			return c.JSON(r.Code, r.Result)
		}

		if strings.ToLower(role.(string)) != "uploader" &&
			strings.ToLower(role.(string)) != "admin" ||
			strings.ToLower(role.(string)) != "administrator" {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for administrator and uploader",
			})
		}

		return next(c)
	}
}
