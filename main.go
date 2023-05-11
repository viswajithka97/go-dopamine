package main

import (
	"go-dopamine/controllers"
	"go-dopamine/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvFile()
	initializers.ConnectDB()
}

func main() {

	r := gin.Default()
	r.GET("/getAllUsers", controllers.GetAllUsers)
	r.GET("/getSuggestedSquads", controllers.GetSuggestedSquads)

	r.Run()
}
