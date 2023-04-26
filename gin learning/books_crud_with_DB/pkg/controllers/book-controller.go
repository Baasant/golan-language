package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"books_crud_DB/pkg/utils"
	"books_crud_DB/pkg/models"
)

 var NewBook models.Book

 func GetbBook(w http.ResponseWriter,r *http.Request)  {
	newBooks:=models.GetAllBooks() // find books and send it back
	res , _=json.Marshal(newBooks)//convert to json
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)//WILL RETURN 200 IF OK
	w.Write(res)
 }//FUNC

 func GetBookById(w http.ResponseWriter,r *http.Request){
	vars :=mux.Vars(r)
    bookId :=vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0) //make sure that the book id be integer instead of string
	if err !=nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_:=models.GetBookById(ID)
	res , _=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)//WILL RETURN 200 IF OK
	w.Write(res)
 }

 func Createbook (w http.ResponseWriter,r *http.Request){
	CreateBook:=models.Book{}
	utils.ParseBody(r,CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
 }

func Deletebook(w http.ResponseWriter,r *http.Request) {
	vars:=mux.vars(r)
	bookId :=vars["bookId"]
	Id,err:=
}