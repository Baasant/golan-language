package main
// download gin package 
//create api 
//struct use in go to keep

// to the communication between server and client must be in json format

import (
	"net/http"
	"github.com/gi-gonic/gin"
)
type todo struct{
	ID   string   'json:"id'
	Item string    'json:"tile'
	completed bool  'json:"completed'
}
//create todo array

var todos=[] todo{
	{ID:"1",Item :"clean room",completed: false},
	{ID:"2",Item :"read book",completed: false},
	{ID:"3",Item :"walking",completed: false},
}

func getTodos(context *gin.context){
	//convert input to json
	context.IndentedJSON(http.)
}
func main(){
	router:=gin.Default()// act as server 
	router.Get("/todos") //
	router.Run("localhost:9090")//path to url
}