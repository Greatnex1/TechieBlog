package main

import (
	"TechieBlog/Initializers"
	"TechieBlog/controllers"
	"github.com/gin-gonic/gin"
)

func init() {

	Initializers.LoadEnvVariables()
	Initializers.ConnectToDB()

}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/all-posts", controllers.PostRead)
	r.GET("/posts/:id", controllers.PostReadById)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/delete/:id", controllers.PostDelete)
	r.Run()
}
