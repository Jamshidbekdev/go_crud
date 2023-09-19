package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamshid/go_crud/initializers"
	routes "github.com/jamshid/go_crud/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	// Create gin app
	app := gin.Default()

	// Configure asset static files
	app.LoadHTMLGlob("templates/views/*")
	app.Static("/assets", "./templates/assets")

	// Routing
	routes.Routes(app)

	// Start app
	app.Run() // listen and serve on 0.0.0.0:8080
}
