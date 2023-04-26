package models

import (
	"books_crud_DB/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:'""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// INTIALIZE database

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) Createbook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books //return all books in the database
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func Deletebook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
