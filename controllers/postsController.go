package controllers

import (
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
	post := models.Post{Title: "Title-1", Body: "Body-1"}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
