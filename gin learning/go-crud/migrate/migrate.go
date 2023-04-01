package main

import (
	"go-crud/initializers"
	"go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
