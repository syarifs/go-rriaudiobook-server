package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/integralist/go-findroot/find"
	"github.com/spf13/viper"
)

var (
	DB_DRIVER                      string
	DB_USERNAME                    string
	DB_PASSWORD                    string
	DB_DATABASE                    string
	DB_HOST                        string
	DB_PORT                        string
	DB_TIMEZONE                    string
	MONGODB_STRING                 string
	MONGODB_DATABASE               string
	SERVER_PORT                    string
	ACCESS_KEY                     []byte
	RESET_KEY                      []byte
	SMTP_SERVER                    string
	SMTP_PORT                      string
	EMAIL                          string
	PASSWORD                       string
	JWT_ACCESS_EXPIRE_TIME         int64
	JWT_REFRESH_EXPIRE_TIME        int64
	JWT_RESET_PASSWORD_EXPIRE_TIME int64
	S3_BUCKET_NAME                 string
	S3_ACCOUNT_ID                  string
	S3_ACCESS_KEY_ID               string
	S3_ACCESS_KEY_SECRET           string
	S3_PUBLIC_ACCESS               string
)

func LoadConfig() {
	var path string
	root, err := find.Repo()
	path = root.Path
	if err != nil {
		path = "."
	}

	viper.AddConfigPath(path + "/config")
	viper.SetConfigName("config")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found, using os env.")
	}

	SERVER_PORT, _ = checkEnv(os.Getenv("PORT"), viper.GetString("server.PORT"))
	MONGODB_STRING, _ = checkEnv(os.Getenv("MONGODB_STRING"), viper.GetString("mongo.STRING"))
	MONGODB_DATABASE, _ = checkEnv(os.Getenv("MONGODB_DATABASE"), viper.GetString("mongo.DATABASE"))
	DB_DRIVER, _ = checkEnv(os.Getenv("DB_DRIVER"), viper.GetString("db.DRIVER"))
	DB_DATABASE, _ = checkEnv(os.Getenv("DB_DATABASE"), viper.GetString("db.DATABASE"))
	DB_USERNAME, _ = checkEnv(os.Getenv("DB_USERNAME"), viper.GetString("db.USERNAME"))
	DB_PASSWORD, _ = checkEnv(os.Getenv("DB_PASSWORD"), viper.GetString("db.PASSWORD"))
	DB_HOST, _ = checkEnv(os.Getenv("DB_HOST"), viper.GetString("db.HOST"))
	DB_PORT, _ = checkEnv(os.Getenv("DB_PORT"), viper.GetString("db.PORT"))
	DB_TIMEZONE, _ = checkEnv(os.Getenv("DB_TIMEZONE"), viper.GetString("db.TIMEZONE"))
	access_key, _ := checkEnv(os.Getenv("ACCESS_KEY"), viper.GetString("secret.ACCESS_KEY"))
	reset_key, _ := checkEnv(os.Getenv("RESET_KEY"), viper.GetString("secret.RESET_KEY"))
	ACCESS_KEY = []byte(access_key)
	RESET_KEY = []byte(reset_key)
	SMTP_SERVER, _ = checkEnv(os.Getenv("SMTP_SERVER"), viper.GetString("smtp.SERVER"))
	SMTP_PORT, _ = checkEnv(os.Getenv("SMTP_PORT"), viper.GetString("smtp.PORT"))
	S3_BUCKET_NAME, _ = checkEnv(os.Getenv("S3_BUCKET_NAME"), viper.GetString("s3.BUCKET_NAME"))
	S3_ACCOUNT_ID, _ = checkEnv(os.Getenv("S3_ACCOUNT_ID"), viper.GetString("s3.ACCOUNT_ID"))
	S3_ACCESS_KEY_ID, _ = checkEnv(os.Getenv("S3_ACCESS_KEY_ID"), viper.GetString("s3.ACCESS_KEY_ID"))
	S3_ACCESS_KEY_SECRET, _ = checkEnv(os.Getenv("S3_ACCESS_KEY_SECRET"), viper.GetString("s3.ACCESS_KEY_SECRET"))
	S3_PUBLIC_ACCESS, _ = checkEnv(os.Getenv("S3_PUBLIC_ACCESS"), viper.GetString("s3.PUBLIC_ACCESS"))
	EMAIL, _ = checkEnv(os.Getenv("SMTP_EMAIL"), viper.GetString("smtp.EMAIL"))
	PASSWORD, _ = checkEnv(os.Getenv("SMTP_PASSWORD"), viper.GetString("smtp.PASSWORD"))
	_, JWT_ACCESS_EXPIRE_TIME = checkEnv(
		os.Getenv("JWT_ACCESS_EXPIRE_TIME"),
		viper.GetString("jwt.ACCESS_TIME"),
		"int",
	)
	_, JWT_REFRESH_EXPIRE_TIME = checkEnv(
		os.Getenv("JWT_REFRESH_EXPIRE_TIME"),
		viper.GetString("jwt.REFRESH_TIME"),
		"int",
	)
	_, JWT_RESET_PASSWORD_EXPIRE_TIME = checkEnv(
		os.Getenv("JWT_RESET_PASSWORD_EXPIRE_TIME"),
		viper.GetString("jwt.FORGOT_PASSWORD_TIME"),
		"int",
	)
}

func checkEnv(args ...interface{}) (envStr string, envInt int64) {
	for i, v := range args {
		switch i {
		case 0:
			if reflect.TypeOf(v).String() == "string" {
				envStr = v.(string)
			} else {
				envInt = v.(int64)
			}
		case 1:
			if envStr != "" || envInt != 0 {
				return
			}

			if reflect.TypeOf(v).String() == "string" {
				envStr = v.(string)
			} else {
				envInt = v.(int64)
			}
		case 2:
			if reflect.TypeOf(v).String() != "string" {
				return
			}

			if v == "int" || v == "int64" {
				intEnv, _ := strconv.Atoi(envStr)
				envInt = int64(intEnv)
				envStr = ""
			}
		}
	}

	return
}
