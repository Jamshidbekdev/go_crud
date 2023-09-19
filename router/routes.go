package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jamshid/go_crud/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/post", controllers.IndexPost)
	router.POST("/post", controllers.CreatePost)
	router.GET("/post/:id", controllers.ShowPost)
	router.PUT("/post/:id", controllers.UpdatePost)
	router.DELETE("/post/:id", controllers.DeletePost)
}
