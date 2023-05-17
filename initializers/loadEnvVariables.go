package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load("/home/ec2-user/go-dopamine")

	if err != nil {
		log.Fatal("Error loading env file")
	}
}
