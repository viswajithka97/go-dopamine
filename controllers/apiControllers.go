package controllers

import (
	"encoding/json"
	"go-dopamine/initializers"
	"go-dopamine/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	// data := []models.GetAllUsers{}

	var user []models.GetAllUsers

	// initializers.
	// 	DB.
	// 	Table("users").
	// 	Joins("left join category_selection ON users.id = category_selection.user_id").
	// 	Select("users.*, GROUP_CONCAT(category_selection.category_id) as category_ids").
	// 	Group("users.id").
	// 	Where("users.id = ?", 585).
	// 	Find(&data)

	initializers.
		DB.Table("users").
		Select("users.*, GROUP_CONCAT(category_selection.category_id SEPARATOR ',') as category_ids").
		Joins("LEFT JOIN category_selection ON users.id = category_selection.user_id").
		Group("users.id").
		Scan(&user)

	// user.CategoryIDs = strings.Split(user.CategoryIDsStr, ",")

	c.JSON(200, gin.H{
		"total":    len(user),
		"response": user,
	})

}

func GetSuggestedSquads(c *gin.Context) {

	// type Squad struct {
	// 	ID   uint   `gorm:"column:squad_id"`
	// 	Name string `gorm:"column:squad_name"`
	// }

	// type Customer struct {
	// 	UserID    uint   `gorm:"column:user_id"`
	// 	FirstName string `gorm:"column:first_name"`
	// }

	type CustomerSquadsResponse struct {
		SquadIDs string `json:"squad_ids"`
	}
	var response CustomerSquadsResponse
	data := []models.SuggestedSquadsModel{}

	var body struct {
		User_Id int `json:"user_id"`
	}

	c.Bind(&body)

	initializers.DB.
		Table("customer").
		Select("squad_members.user_id, JSON_ARRAYAGG(squad_members.squad_id) as squad_ids").
		Joins("LEFT JOIN squad_members ON customer.user_id = squad_members.user_id").
		Joins("LEFT JOIN squads ON squad_members.squad_id = squads.id").
		Where("customer.user_id = ?", body.User_Id).
		Where("JSON_CONTAINS(customer.interests, cast(squads.category as char), '$')").
		Group("customer.user_id").
		Find(&response)

	log.Println(response.SquadIDs)

	var squadIds []int
	json.Unmarshal([]byte(response.SquadIDs), &squadIds)

	log.Println(squadIds)

	initializers.DB.Table("squads").Where("squads.id = ?", squadIds).Find(&data)

	c.JSON(200, gin.H{
		"response": data,
	})

}

func GetAllExperts(c *gin.Context) {
	
type Sample struct{
	name string
	age int64
	email string
}

data := Sample{
	name : "Ameer",
	age : 26,
	email : "ameer@gmail.com",
}

log.Println(data)

c.JSON(200, gin.H{
	"response" : data,
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

