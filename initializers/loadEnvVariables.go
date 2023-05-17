package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load("/home/ec2-user/dopamine/go-dopamine/go-dopamine/.env")

	if err != nil {
		log.Fatal("Error loading env file")
	}
}
