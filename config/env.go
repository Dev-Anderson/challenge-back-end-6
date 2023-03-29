package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func LoadEnv() (Env, error) {
	godotenv.Load()
	return Env{
		Host:     os.Getenv("ADOPET_HOST"),
		Port:     os.Getenv("ADOPET_PORT"),
		User:     os.Getenv("ADOPET_USER"),
		Password: os.Getenv("ADOPET_PASSWORD"),
		Database: os.Getenv("ADOPET_DATABASE"),
	}, nil
}
