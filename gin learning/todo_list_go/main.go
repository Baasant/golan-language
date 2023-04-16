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

	//this struct we need it for the front end 
	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreateAt  time.Time `json:"created_at"`
	}
)
///////////////////////////

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
// R ITS A POINTER TO HTTP REQUEST
func homeHandler(w http.ResponseWriter,r*http.Request){
	//if someone goes to my localhost it will rander the static
err:=rnd.Template(w,hhtp/statusOk,[]string{"static/home.tpl"},nil)
checkErr(err)
}

func fetchTodos(w http.ResponseWriter,r *http.Request){
	todo:=[]todoModel{}
	if err :=db.c(collectionName).Find(bson.M{}).All(&todos); err!=nil{
		rnd.JSON(w,http.StatusProcessing,renderer.M{
			"message":"Failed to fetch todo",
			"error":err,
		})
		return
	}
	todoList :=[]todos{} //will be json

	for _,t:=range todos{
		todoList=append(toList,todo{
			ID:t.ID.Hex(),
			Tilte:t.t.Title,
			Completed:t.Completed,
			CreatedAt:t.CreateAt,
		})
	}
	rnd.JSON(w,hhtp.statusOk,renderer.M{
		"data":todoList,
	})
}
//define main func

func main (){
	//create router 
	//use to stop code 
	stopShan :=make(chan os.Signal)
	signal.Notify(stopChan,os.Interrupt)
	r:=chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/",homeHandler)
	r.Mount("/todo",todoHandlers())
	srv:=&http.Server{
		Addr:port,
		Handler:r,
		ReadTimeout:  60*time.Second,
		WriteTimeout: 60*time.Second,
		IdleTimeout :60*time.Second,
	}
	//////////used to stop code///////////////
	go func() {
		log.Println("Listening on port ", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
	////////////////////////////
}

<-stopChan
log.Println()

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
