package config

import "github.com/spf13/viper"

var AppConfig *Config

type Config struct {
	*viper.Viper
}

func LoadAllConfig() {
	loadDbConfig()
	loadApp()
}
