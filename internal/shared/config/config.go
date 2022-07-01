package config

import (
	"github.com/joho/godotenv"
	"os"
)

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Schema   string
}

type Config struct {
	DB         ConfigDB
	ListenPort string
}

var config *Config

func InstanceConfig() Config {
	if config == nil {
		initEnvironment()
	}
	return *config
}

func initEnvironment() {
	_ = godotenv.Load(os.Getenv("ENV_FILE"))
	config = &Config{
		DB: ConfigDB{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Schema:   os.Getenv("DB_SCHEMA"),
			Host:     os.Getenv("DB_HOST"),
		},
		ListenPort: os.Getenv("SERVICE_PORT"),
	}
}
