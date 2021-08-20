package util

import "github.com/spf13/viper"

type Config struct {
	DbDriver                  string `mapstructure: "DB_DRIVER"`
	DbSource                  string `mapstructure: "DB_SOURCE"`
	DbMaxConnections          int    `mapstructure: "DB_MAX_CONNECTIONS"`
	DbMaxIdleConnections      int    `mapstructure: "DB_MAX_IDLE_CONNECTIONS"`
	DbMaxLifetimeConnections  int    `mapstructure: "DB_MAX_LIFETIME_CONNECTIONS"`
	ServerUrl                 string `mapstructure: "SERVER_URL"`
	ServerReadTimeout         int    `mapstructure: "SERVER_READ_TIMEOUT"`
	JWTKey                    string `mapstructure: "JWT_SECRET_KEY"`
	JWTKeyExpirationInMinutes int    `mapstructure: "JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"`
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
