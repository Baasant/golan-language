package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data of request body
	// create a post
	post := models.Post{Title: "hello", Body: "post body"}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	// check if there is error
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
		// "message":  "keep going Bassant,i love you ,you are doing great job",
		// "message2": "here we used elepentsql",
		// "message3": "using pstagesql local and tablePlus",
	})
}
