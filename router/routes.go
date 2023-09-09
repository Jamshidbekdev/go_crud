package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jamshid/go_crud/controllers"
)

func Routes() {
	router := gin.Default()
	router.GET("/", controllers.Index)
	router.POST("/post", controllers.CreatePost)

	router.Run() // listen and serve on 0.0.0.0:8080
}
