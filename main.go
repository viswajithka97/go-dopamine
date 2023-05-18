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
	r.POST("/getSuggestedSquads", controllers.GetSuggestedSquads)
	r.GET("/getAllExperts", controllers.GetAllExperts)
	r.GET("/getAllCustomers", controllers.GetAllCustomers)

	r.Run()
}
