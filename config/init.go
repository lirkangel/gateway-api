package config

import (
	"github.com/spf13/viper"
)

type CONFIG struct {
	APP_NAME string
	APP_PORT int
}

func Init() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName(".env")

	return viper.ReadInConfig()
}
