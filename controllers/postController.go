package controllers

import (
	"TechieBlog/Initializers"
	"TechieBlog/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := Initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func PostRead(c *gin.Context) {
	var post []models.Post
	Initializers.DB.Find(&post)

	c.JSON(200, gin.H{"post": post})
}
func PostReadById(c *gin.Context) {

	id := c.Param("id")
	var post models.Post
	Initializers.DB.Find(&post, id)

	c.JSON(200, gin.H{"post": post})
}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	var post models.Post
	Initializers.DB.Find(&post, id)

	Initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{"post": post})
}

func PostDelete(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	Initializers.DB.Delete(&post, id)

	c.Status(200)

}
