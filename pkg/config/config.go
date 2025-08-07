package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	*viper.Viper
}

// New : return new config instance
func NewConfiguration(env string) *Config {
	log.Printf("Init app config for env: %s", env)
	config := &Config{
		Viper: viper.New(),
	}

	config.setDefault()

	return config
}

func (config *Config) setDefault() {
	config.SetDefault("MIDDLEWARE_GIN_RECOVER_ENABLED", true)
	config.SetDefault("MIDDLEWARE_GIN_LOGGING_ENABLED", true)
}

func LoadAllConfigServer() {
	loadDbConfig()
	loadApp()
}

func LoadAllConfigGrpc() {
	loadDbConfig()
	loadGrpc()
}
