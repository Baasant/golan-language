package controllers

import (
	"books_crud_with_DB/pkg/models"
	"books_crud_with_DB/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() // find books and send it back
	res, _ := json.Marshal(newBooks) //convert to json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //WILL RETURN 200 IF OK
	w.Write(res)
} //FUNC

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) //make sure that the book id be integer instead of string
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //WILL RETURN 200 IF OK
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)      //access the request that we recieve from user
	bookId := vars["bookId"] //acess book id
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.Deletebook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //WILL RETURN 200 IF OK
	w.Write(res)
}

// update function
// we you nwant to update a book ,you need to have the id of that book
// and the new updatred book information such as name ,publication
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//this function is nix between create and getbyid functions
	///////like createbook func//////////
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	/////////////////////////////////
	////////getbookbyid function
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	//get book by id
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	} //if update book
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	} //if update book
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	} //if update book
	////////////////////////////////////
	//save new data in the database
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //WILL RETURN 200 IF OK
	w.Write(res)
} //function
