package main

import (
	"nothing.example/initializers"
	"nothing.example/models"
)

//	func init() {
//		initializers.LoadEnvVariables()
//		initializers.ConnectToDB()
//	}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
