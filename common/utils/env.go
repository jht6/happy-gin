package utils

import (
	"os"
)

// SERVER_ENV可选值: prod dev
func GetServerEnv() string {
	env := os.Getenv("SERVER_ENV")
	if env == "" {
		env = "prod"
	}
	return env
}

func IsProd() bool {
	return GetServerEnv() == "prod"
}

func IsDev() bool {
	return GetServerEnv() == "dev"
}
