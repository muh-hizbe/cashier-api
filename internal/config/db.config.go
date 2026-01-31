package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	URL string
}

var dbConfig DBConfig

func GetDBConfig() DBConfig {
	return dbConfig
}

func setDefaultDBConfig() {
	viper.SetDefault("DB_URL", "")
}

func SetDBConfig(config DBConfig) {
	dbConfig = config
}
