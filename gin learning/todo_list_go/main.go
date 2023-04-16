package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("!... Hello World ...!")
// }

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render

var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = ":todo"
	port           string = ":9000"
)

//crreate todo model that we will use with mangodb database

// we need json to interact with db and another one to interact with front end
type (
	todoModel struct {
		ID        bson.ObjectId `bosn:_"id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `"bson:"completed"`
		CreateAt  time.Time     `"bson:"createAt"`
	}

	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreateAt  time.Time `json:"created_at"`
	}
)

// conect with db and start  the session
func init() {
	rnd = renderer.New()
	sess, err :=
		mgo.Database(hostName)

	//check if there is an error
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

//define main func

func main (){
	//create router 
	r:=chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/",homeHandler)
	r.Mount("/todo",todoHandlers())
	srv:=&http.Server{
		Addr:port,
		Handler:r,
		ReadTimeout: : 60*time.Second,
		WriteTimeout: 60*time.Second,
		IdleTimeout :60*time.Second,
	}
	go func (){
		log.Println("Listening on port",port)
		if err:srv.ListenAndServe(); err! nil {
			log.Printf("listen%s\n",err)
		}
	}
}

func todoHandlers() http.Handler{
	rg:=CHI.NewRouter()
	rg.Group(func (r chi.Router){
		r.Get("/",fetchTodos)
		r.Post("/",createTodo)
		r.Put("/{id}",updateTodo)
		r.Delete("/{id}",deleteTodo)

	})
return rg	 
}
// if there is error log out 
func checkErr(err error)
{
if err!nil{
	log.Fatal(err)
}
}
