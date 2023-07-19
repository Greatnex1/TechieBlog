package main

import (
	"TechieBlog/Initializers"
	"TechieBlog/models"
	"log"
)

func init() {
	Initializers.LoadEnvVariables()
	Initializers.ConnectToDB()
}

func main() {
	err := Initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Migration failed")
	}
}
