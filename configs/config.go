package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost string
	DbName string
	DbUser string
	DbPass string
}

var config *Config

func Load() *Config {
	if config == nil {
		if err := godotenv.Load(); err != nil {
			panic(err.Error())
		}
		config = &Config{
			DbUser: os.Getenv("DB_USERNAME"),
			DbPass: os.Getenv("DB_PASSWORD"),
			DbHost: os.Getenv("DB_HOST"),
			DbName: os.Getenv("DB_NAME"),
		}
	}

	return config
}
