package controllers

import (
	"go-dopamine/initializers"
	"go-dopamine/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	data := []models.GetAllUsers{}

	initializers.
		DB.
		Table("users").
		Joins("left join category_selection ON users.id = category_selection.user_id").
		Select("users.*, GROUP_CONCAT(category_selection.category_id) as category_ids").
		Group("users.id").
		// Where("users.id = ?", 585).
		Find(&data)

	c.JSON(200, gin.H{
		// "total":    len(data),
		"response": data,
	})

}

// func FetchAllPosts(c *gin.Context) {

// 	// get the posts

// 	var posts []models.Post

// 	initializers.DB.Find(&posts)
// 	// response

// 	c.JSON(200, gin.H{
// 		"posts": posts,
// 	})

// }

// func PostById(c *gin.Context) {

// 	// get id from params

// 	id := c.Param("id")

// 	// get the posts

// 	var post models.Post

// 	initializers.DB.First(&post, id)
// 	// response

// 	c.JSON(200, gin.H{
// 		"posts": post,
// 	})

// }

// func PostUpdate(c *gin.Context) {

// 	// get id from params

// 	id := c.Param("id")

// 	// get the data of the req body

// 	var body struct {
// 		Title string
// 		Body  string
// 	}

// 	c.Bind(&body)

// 	//find the post we are updating

// 	var post models.Post

// 	initializers.DB.First(&post, id)

// 	// update it

// 	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

// 	// send response

// 	c.JSON(200, gin.H{
// 		"post": post,
// 	})

// }
// func PostDelete(c *gin.Context) {

// 	// get id from params

// 	id := c.Param("id")

// 	// delete it

// 	initializers.DB.Delete(&models.Post{}, id)

// 	// send response

// 	c.JSON(200, gin.H{
// 		"message": "post deleted",
// 	})

// }

// type expertData struct {
// 	ID        int64  `json:"id"`
// 	User_ID   int64  `json:"user_id"`
// 	FirstName string `json:"first_name"`
// 	 gorm.Model
// 	 ID        uint
// 	 FirstName string
// 	 UserID    uint
// }

// expert := expertData{}

// initializers.DB.Table("expert").Joins("LEFT JOIN users ON users.id = expert.user_id").Where("users.id = ?", 585).Find(&expert)
// log.Println(expert)
