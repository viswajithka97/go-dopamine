package initializers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envFilePath := filepath.Join(dir, ".env")

	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// err := godotenv.Load("/home/ec2-user/dopamine/go-dopamine/go-dopamine/.env")

	// if err != nil {
	// 	log.Fatal("Error loading env file")
	// }
}
