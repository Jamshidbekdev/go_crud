package main

import (
	"github.com/jamshid/go_crud/initializers"
	routes "github.com/jamshid/go_crud/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	routes.Routes()
}
