package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/integralist/go-findroot/find"
	"github.com/spf13/viper"
)

var (
	DB_DRIVER                        string
	DB_USERNAME                      string
	DB_PASSWORD                      string
	DB_DATABASE                      string
	DB_HOST                          string
	DB_PORT                          string
	DB_TIMEZONE                      string
	MONGODB_STRING                   string
	MONGODB_DATABASE                 string
	SERVER_PORT                      string
	SMTP_SERVER                      string
	SMTP_PORT                        string
	EMAIL                            string
	PASSWORD                         string
	TOKEN_DRIVER                     string
	TOKEN_ACCESS_EXPIRE_TIME         int64
	TOKEN_REFRESH_EXPIRE_TIME        int64
	TOKEN_RESET_PASSWORD_EXPIRE_TIME int64
	ACCESS_KEY                       []byte
	RESET_KEY                        []byte
	S3_BUCKET_NAME                   string
	S3_ACCOUNT_ID                    string
	S3_ACCESS_KEY_ID                 string
	S3_ACCESS_KEY_SECRET             string
	S3_PUBLIC_ACCESS                 string
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

	SERVER_PORT = stringEnv(os.Getenv("PORT"), viper.GetString("server.PORT"))
	MONGODB_STRING = stringEnv(os.Getenv("MONGODB_STRING"), viper.GetString("mongo.STRING"))
	MONGODB_DATABASE = stringEnv(os.Getenv("MONGODB_DATABASE"), viper.GetString("mongo.DATABASE"))
	DB_DRIVER = stringEnv(os.Getenv("DB_DRIVER"), viper.GetString("db.DRIVER"))
	DB_DATABASE = stringEnv(os.Getenv("DB_DATABASE"), viper.GetString("db.DATABASE"))
	DB_USERNAME = stringEnv(os.Getenv("DB_USERNAME"), viper.GetString("db.USERNAME"))
	DB_PASSWORD = stringEnv(os.Getenv("DB_PASSWORD"), viper.GetString("db.PASSWORD"))
	DB_HOST = stringEnv(os.Getenv("DB_HOST"), viper.GetString("db.HOST"))
	DB_PORT = stringEnv(os.Getenv("DB_PORT"), viper.GetString("db.PORT"))
	DB_TIMEZONE = stringEnv(os.Getenv("DB_TIMEZONE"), viper.GetString("db.TIMEZONE"))
	access_key := stringEnv(os.Getenv("ACCESS_KEY"), viper.GetString("secret.ACCESS_KEY"))
	reset_key := stringEnv(os.Getenv("RESET_KEY"), viper.GetString("secret.RESET_KEY"))
	ACCESS_KEY = []byte(access_key)
	RESET_KEY = []byte(reset_key)
	SMTP_SERVER = stringEnv(os.Getenv("SMTP_SERVER"), viper.GetString("smtp.SERVER"))
	SMTP_PORT = stringEnv(os.Getenv("SMTP_PORT"), viper.GetString("smtp.PORT"))
	S3_BUCKET_NAME = stringEnv(os.Getenv("S3_BUCKET_NAME"), viper.GetString("s3.BUCKET_NAME"))
	S3_ACCOUNT_ID = stringEnv(os.Getenv("S3_ACCOUNT_ID"), viper.GetString("s3.ACCOUNT_ID"))
	S3_ACCESS_KEY_ID = stringEnv(os.Getenv("S3_ACCESS_KEY_ID"), viper.GetString("s3.ACCESS_KEY_ID"))
	S3_ACCESS_KEY_SECRET = stringEnv(os.Getenv("S3_ACCESS_KEY_SECRET"), viper.GetString("s3.ACCESS_KEY_SECRET"))
	S3_PUBLIC_ACCESS = stringEnv(os.Getenv("S3_PUBLIC_ACCESS"), viper.GetString("s3.PUBLIC_ACCESS"))
	EMAIL = stringEnv(os.Getenv("SMTP_EMAIL"), viper.GetString("smtp.EMAIL"))
	PASSWORD = stringEnv(os.Getenv("SMTP_PASSWORD"), viper.GetString("smtp.PASSWORD"))
	TOKEN_DRIVER = stringEnv(os.Getenv("TOKEN_DRIVER"), viper.GetString("token.DRIVER"))
	TOKEN_ACCESS_EXPIRE_TIME = intEnv(
		os.Getenv("TOKEN_ACCESS_EXPIRE_TIME"),
		viper.GetString("token.ACCESS_TIME"),
	)
	TOKEN_REFRESH_EXPIRE_TIME = intEnv(
		os.Getenv("TOKEN_REFRESH_EXPIRE_TIME"),
		viper.GetString("token.REFRESH_TIME"),
	)
	TOKEN_RESET_PASSWORD_EXPIRE_TIME = intEnv(
		os.Getenv("TOKEN_RESET_PASSWORD_EXPIRE_TIME"),
		viper.GetString("token.FORGOT_PASSWORD_TIME"),
	)
}

func stringEnv(os, viper string) (env string) {
	if os != "" {
		return os
	}
	return viper
}

func intEnv(os, viper string) (env int64) {
	if os != "" {
		int, _ := strconv.Atoi(os)
		env = int64(int)
		return
	}
	int, _ := strconv.Atoi(viper)
	env = int64(int)
	return
}
