package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER                           string
	DB_SOURCE                           string
	DB_MAX_CONNECTIONS                  int
	DB_MAX_IDLE_CONNECTIONS             int
	DB_MAX_LIFETIME_CONNECTIONS         int
	SERVER_URL                          string
	SERVER_READ_TIMEOUT                 int
	JWT_KEY                             string
	JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT time.Duration
	TOKEN_SYMMETRIC_KEY                 string
	PASETO_TOKEN_DURATION               time.Duration
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
