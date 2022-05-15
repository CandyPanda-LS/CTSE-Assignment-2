package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURL() string {
	godotenv.Load()

	return os.Getenv("MONGO_URL")
}

func EnvEmailService() string {
	godotenv.Load()

	return os.Getenv("EMAIL_SERVICE")
}
