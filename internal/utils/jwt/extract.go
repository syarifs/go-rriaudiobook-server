package jwt

import (
	"context"
	"fmt"
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/utils/config"
	"go-rriaudiobook-server/internal/utils/errors"
	"go-rriaudiobook-server/internal/utils/logger"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var mongoc *mongo.Database
var sqlc *gorm.DB

func NewJWTConnection(mongo *mongo.Database, sql *gorm.DB) {
	mongoc = mongo
	sqlc = sql
}

func GetTokenData(header string, data string, tokenType Token) (extracted interface{}, err error) {
	err = FindToken(header)
	if err != nil {
		return
	}

	extract, err := ExtractToken(header, tokenType)
	if extract != nil {
		extracted, _ = extract.(jwt.MapClaims)[data]
	}

	return
}

func GetToken(c echo.Context, tokenType Token) (header string, err error) {
	header = c.Request().Header.Get("Authorization")
	headers := strings.Split(header, " ")

	if header == "" || len(headers) < 2 {
		err = errors.ErrNoToken
		return
	}

	header = headers[1]
	_, err = ExtractToken(header, tokenType)
	return
}

func ExtractToken(tkn string, tokenType Token) (token interface{}, err error) {
	token, err = jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		var key []byte

		if tokenType == ACCESS {
			key = config.ACCESS_KEY
		} else if tokenType == RESET {
			key = config.RESET_KEY
		}

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})

	_token := token.(*jwt.Token)
	if !_token.Valid {
		return nil, errors.ErrInvalidToken
	}

	if err != nil {
		return nil, err
	}

	return token.(*jwt.Token).Claims, nil
}

func FindToken(token string) error {
	var mongoData []bson.M
	var sqlData models.Token

	if mongoc != nil {
		filter := bson.D{
			{Key: "access_token", Value: token},
		}

		db := mongoc.Collection("token")
		cur, err := db.Find(context.TODO(), filter)

		if err != nil {
			logger.WriteLog(err)
		}

		if err = cur.All(context.TODO(), &mongoData); err != nil {
			logger.WriteLog(err)
		}

	} else if sqlc != nil {
		sqlc.Model(models.Token{}).
			Where("access_token = ?", token).Scan(&sqlData)
	}

	if mongoData == nil && sqlData.AccessToken == "" {
		return errors.ErrInvalidToken
	}

	return nil
}
