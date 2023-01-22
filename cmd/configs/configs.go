package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetMongoURL() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	githubUsername, exists := os.LookupEnv("MONGO_URL")

	if exists {
		fmt.Println(githubUsername)
	}
}
