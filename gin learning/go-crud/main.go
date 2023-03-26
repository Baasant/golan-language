package main

// import "github.com/gin-gonic/gin"
import (
	"github.com/gin-gonic/gin"
	// "github.com/robbyklein/go-crud/initializers"
	"go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "keep going Bassant,i love you ,you are doing great job",
			"message2": "here we used elepentsql",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
