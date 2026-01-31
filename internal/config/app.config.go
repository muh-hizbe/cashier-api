package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port    string
	Version string
}

var appConfig AppConfig

func GetAppConfig() AppConfig {
	return appConfig
}

func setDefaultAppConfig() {
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("APP_VERSION", "1.0.0")
}

func SetAppConfig(config AppConfig) {
	appConfig = config
}
