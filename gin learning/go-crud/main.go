package main

// import "github.com/gin-gonic/gin"
import (
	"github.com/gin-gonic/gin"
	// "github.com/robbyklein/go-crud/initializers"
	"go-crud/controllers"
	"go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()

}

func main() {
	r := gin.Default()
	// connect every things to our router
	// r.GET("/", controllers.PostsCreate)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdates)
	r.DELETE("/posts/:id", controllers.PostsDlete)

	r.Run() // listen and serve on 0.0.0.0:8080 the we change it to run on 3000
}
