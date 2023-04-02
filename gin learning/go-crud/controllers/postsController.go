package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off req body

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// get data of request body
	// create a post
	// post := models.Post{Title: "hello bassant", Body: "post body"}
	post := models.Post{Title: body.Title, Body: body.Body}
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

func PostsIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	// find all the posts and assign it to the variable posts
	initializers.DB.Find(&posts)
	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
		// "message":  "keep going Bassant,i love you ,you are doing great job",
		// "message2": "here we used elepentsql",
		// "message3": "using pstagesql local and tablePlus",
	})
}

// function to fetch single post
func PostsShow(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//get the posts
	var post models.Post
	// find all the posts and assign it to the variable posts
	initializers.DB.First(&post, id)
	//respond with them

	c.JSON(200, gin.H{
		"post": post,
		// "message":  "keep going Bassant,i love you ,you are doing great job",
		// "message2": "here we used elepentsql",
		// "message3": "using pstagesql local and tablePlus",
	})
}

func PostsUpdates(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//get the data of request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// find the post you want to  updated
	var post models.Post
	// find all the posts and assign it to the variable posts
	initializers.DB.First(&post, id)
	//Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title,
		Body: body.Body,
	})

	//response with it
	c.JSON(200, gin.H{
		"post": post,
		// "message":  "keep going Bassant,i love you ,you are doing great job",
		// "message2": "here we used elepentsql",
		// "message3": "using pstagesql local and tablePlus",
	})
}

func PostsDlete(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	// delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//response with it
	// get 200 if its deleted correctly
	c.Status(200)

}
