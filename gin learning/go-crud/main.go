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
	// r.GET("/", controllers.PostsCreate)
	r.POST("/posts", controllers.PostsCreate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
