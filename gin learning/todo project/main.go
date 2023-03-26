package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = strconv.Itoa(len(todos) + 1)
	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if todo.ID == id {
			var updatedTodo Todo

			if err := c.ShouldBindJSON(&updatedTodo); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			updatedTodo.ID = todo.ID
			todos[i] = updatedTodo

			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
