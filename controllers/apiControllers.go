package controllers

import (
	"encoding/json"
	"fmt"
	"go-dopamine/initializers"
	"go-dopamine/models"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	// data := []models.GetAllUsers{}

	var data []models.GetAllUsers

	initializers.
		DB.Table("users").
		Select("users.*, GROUP_CONCAT(category_selection.category_id SEPARATOR ',') as category_ids").
		Joins("LEFT JOIN category_selection ON users.id = category_selection.user_id").
		Group("users.id").
		Scan(&data)

	// for i := range data {

	// 	data[i].CategoryIds =
	// 		strings.Split(data[i].CategoryIDsStr, ",")
	// }

	for i := range data {
		if data[i].CategoryIDsStr != "" {
			categoryIDsStr := strings.Split(data[i].CategoryIDsStr, ",")
			categoryIDs := make([]int, len(categoryIDsStr))

			for j, str := range categoryIDsStr {
				num, err := strconv.Atoi(str)
				if err != nil {
					// Handle conversion error
					fmt.Printf("Error converting string to int: %v\n", err)
					continue
				}
				categoryIDs[j] = num
			}

			data[i].CategoryIds = categoryIDs
		}

	}

	// user.CategoryIds = strings.Split(user.CategoryIDsStr, ",")

	c.JSON(200, gin.H{
		"total":    len(data),
		"response": data,
		
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
	
// type Sample struct{
// 	Name string
// 	Age int64
// 	Email string
// }

// data := Sample{
// 	Name : "Ameer",
// 	Age : 26,
// 	Email : "ameer@gmail.com",
// }

// log.Println(data)

var experts []models.ExpertListModel

initializers.
	DB.Table("expert e").
	Select("e.id, e.first_name, e.category_ids").
	Scan(&experts)

	fmt.Printf("initial %T\n",experts[0].Category_ids)

	for i := 0; i < len(experts); i++ {
		strings.Trim(experts[i].Category_ids, "[]")
		strings.Split(experts[i].Category_ids, ",")
		strconv.ParseInt(experts[i].Category_ids, 10, 64)
	}

	fmt.Printf("converted %T\n",experts[0].Category_ids)

c.JSON(200, gin.H{
	"response" : experts,
})

}

func GetAllCustomers(c *gin.Context) {

	var user []models.GetAllUsers

	initializers.DB.Table("customer").Scan(&user)

	var isTrue bool

	if len(user) == 0 {
		isTrue = false
	} else {
		isTrue = true
	}

	if len(user) == 0 {
		c.JSON(200, gin.H{
			"success":  isTrue,
			"response": "Customers are empty",
			"meta":     len(user),
		})

	} else {
		c.JSON(200, gin.H{
			"success":  isTrue,
			"response": user,
			"meta":     len(user),
		})
	}

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
