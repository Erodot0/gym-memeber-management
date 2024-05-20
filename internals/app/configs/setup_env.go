package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func InitializeEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}