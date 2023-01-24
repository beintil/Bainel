package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetMongoURL() string {
	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	login, _ := os.LookupEnv("MONGO_URL")

	return login
}
