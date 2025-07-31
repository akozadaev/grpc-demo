package helper

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	envPath := "."
	envFileName := ".env"

	fullPath := envPath + "/" + envFileName

	if err := godotenv.Overload(fullPath); err != nil {
		log.Printf("[ERROR] failed with %+v", "No .env file found")
	}

	return os.Getenv(key)
}
