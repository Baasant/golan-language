package main

// import "github.com/gin-gonic/gin"

import (
	"go-jwt/initializers"

	"go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDb()
	initializers.SyncDatabase()

}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.Run()
}
