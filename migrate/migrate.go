package main

import (
	"github.com/madhu0309/go-crud/initializers"
	"github.com/madhu0309/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
