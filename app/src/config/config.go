package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMySQLUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		GetEnv("DB_USER"),
		GetEnv("DB_PASS"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_NAME"),
	)
}

func GetEnv(key string) string {

	if os.Getenv(key) == "" {
		log.Fatalf("Environment variable %s not set.", key)
	}

	return os.Getenv(key)
}
