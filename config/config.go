package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string, defaultValue ...string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("An error occurred when loading .env file!")
	}

	value, present := os.LookupEnv(key)

	if present != true && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return value
}
