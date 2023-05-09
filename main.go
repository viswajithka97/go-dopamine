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
	r.GET("/getAllPosts", controllers.FetchAllPosts)
	r.GET("/postById/:id", controllers.PostById)
	r.PUT("/updatePost/:id", controllers.PostUpdate)
	r.DELETE("/deletePost/:id", controllers.PostDelete)

	r.Run()
}
