package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DbType string
	DbLink string
}

func NewConfig() *Config {

	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	dbLink := os.Getenv("DB_LINK")
	if !Validate(dbLink) {
		panic("DB_LINK is not set")
	}

	dbType := os.Getenv("DB_TYPE")
	if !Validate(dbType) {
		panic("DB_TYPE is not set")
	}

	return &Config{
		DbType: dbType,
		DbLink: dbLink,
	}
}

func Validate(param string) bool {
	return param != ""
}
