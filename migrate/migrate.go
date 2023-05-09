package main

import (
	"go-dopamine/initializers"
	"go-dopamine/models"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
