package config

import (
	"fmt"
	"os"

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

	SERVER_PORT = checkStringEnv(os.Getenv("PORT"), viper.GetString("server.PORT"))
	MONGODB_STRING = checkStringEnv(os.Getenv("MONGODB_STRING"), viper.GetString("mongo.STRING"))
	MONGODB_DATABASE = checkStringEnv(os.Getenv("MONGODB_DATABASE"), viper.GetString("mongo.MONGODB_DATABASE"))
	DB_DRIVER = checkStringEnv(os.Getenv("DB_DRIVER"), viper.GetString("db.DRIVER"))
	DB_DATABASE = checkStringEnv(os.Getenv("DB_DATABASE"), viper.GetString("db.DATABASE"))
	DB_USERNAME = checkStringEnv(os.Getenv("DB_USERNAME"), viper.GetString("db.USERNAME"))
	DB_PASSWORD = checkStringEnv(os.Getenv("DB_PASSWORD"), viper.GetString("db.PASSWORD"))
	DB_HOST = checkStringEnv(os.Getenv("DB_HOST"), viper.GetString("db.HOST"))
	DB_PORT = checkStringEnv(os.Getenv("DB_PORT"), viper.GetString("db.PORT"))
	DB_TIMEZONE = checkStringEnv(os.Getenv("DB_TIMEZONE"), viper.GetString("db.TIMEZONE"))
	ACCESS_KEY = []byte(checkStringEnv(os.Getenv("ACCESS_KEY"), viper.GetString("secret.ACCESS_KEY")))
	RESET_KEY = []byte(checkStringEnv(os.Getenv("RESET_KEY"), viper.GetString("secret.RESET_KEY")))
	SMTP_SERVER = checkStringEnv(os.Getenv("SMTP_SERVER"), viper.GetString("smtp.SERVER"))
	SMTP_PORT = checkStringEnv(os.Getenv("SMTP_PORT"), viper.GetString("smtp.PORT"))
	EMAIL = checkStringEnv(os.Getenv("EMAIL"), viper.GetString("smtp.EMAIL"))
	PASSWORD = checkStringEnv(os.Getenv("PASSWORD"), viper.GetString("smtp.PASSWORD"))
	JWT_ACCESS_EXPIRE_TIME = checkInt64Env(os.Getenv("JWT_ACCESS_EXPIRE_TIME"), viper.GetInt64("jwt.ACCESS_TIME"))
	JWT_REFRESH_EXPIRE_TIME = checkInt64Env(os.Getenv("JWT_REFRESH_EXPIRE_TIME"), viper.GetInt64("jwt.REFRESH_TIME"))
	JWT_RESET_PASSWORD_EXPIRE_TIME = checkInt64Env(os.Getenv("JWT_RESET_PASSWORD_EXPIRE_TIME"), viper.GetInt64("jwt.FORGOT_PASSWORD_TIME"))
}

func checkStringEnv(env, viperStr interface{}) string {
	if env != "" {
		return env.(string)
	}
	return viperStr.(string)
}

func checkInt64Env(env, viperStr interface{}) int64 {
	if env != "" {
		return env.(int64)
	}
	return viperStr.(int64)
}
