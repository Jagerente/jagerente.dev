package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host string `mapstructure:"HOST"`
}

var ENV Config

func Init() (err error) {
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		return err
	}

	return nil
}
