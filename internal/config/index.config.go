package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DBConfig
}

func LoadConfig() Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	setupConfig()
	return GetConfig()
}

func GetConfig() Config {
	return Config{
		Database: GetDBConfig(),
		App:      GetAppConfig(),
	}
}

func setupConfig() {
	setDefaultConfig()

	SetAppConfig(AppConfig{
		Version: viper.GetString("APP_VERSION"),
		Port:    viper.GetString("APP_PORT"),
	})

	SetDBConfig(DBConfig{
		URL: viper.GetString("DB_URL"),
	})
}

func setDefaultConfig() {
	setDefaultAppConfig()
	setDefaultDBConfig()
}
