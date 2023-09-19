package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamshid/go_crud/initializers"
	"github.com/jamshid/go_crud/models"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreatePost(c *gin.Context) {

	// create body struct
	var body struct {
		Title string
		Body  string
	}

	//bind all data to body struct
	c.Bind(&body)

	// make Post object and fill it with given data
	post := models.Post{Title: body.Title, Body: body.Body}

	// create instance in db
	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return response post data with status 200
	c.JSON(200, gin.H{
		"post": post,
	})
}

func IndexPost(c *gin.Context) {

	// create array consist of instance of models.Post struct
	var posts []models.Post

	//get all post data
	initializers.DB.Find(&posts) // pass pointer of data to get

	c.HTML(http.StatusOK, "post.tmpl", gin.H{
		"posts": posts,
	})

	// // return response post data with status 200
	// c.JSON(200, gin.H{
	// 	"posts": posts,
	// })
}

func ShowPost(c *gin.Context) {

	// get id param from request
	id := c.Param("id")

	// create instance of models.Post struct
	var post models.Post

	// get one Post data by id
	initializers.DB.First(&post, id) // pass pointer of data to get

	// return response post data with status 200
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {

	// get id param from request
	id := c.Param("id")

	// create body struct
	var body struct {
		Body  string
		Title string
	}

	//bind all data to body struct
	c.Bind(&body)

	// create instance of models.Post struct
	var post models.Post

	// get one Post data by id
	initializers.DB.First(&post, id)

	// update this instance with new data
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// return response post data with status 200
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// get id param from request
	id := c.Param("id")

	// delete instane in posts

	initializers.DB.Delete(&models.Post{}, id)

	// return response post data with status 200
	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})
}
