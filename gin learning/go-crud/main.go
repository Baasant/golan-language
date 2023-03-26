package main

// import "github.com/gin-gonic/gin"
import (
	"github.com/gin-gonic/gin"
	"github.com/robbyklein/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "keep going",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
