package main

import (
	"github.com/jamshid/go_crud/initializers"
	"github.com/jamshid/go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
