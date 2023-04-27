// the routes will hit this routes
package routes

import (
	"books_crud_with_DB/pkg/controllers" //to be able to acess files inside controllers

	"github.com/gorilla/mux"
)

// use this function to create handlers
// var RegisterBookStoresRouters = func(router *mux.Router) {
// 	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")         //if someone come to my/book/ send him/her the createbook func to handel it
// 	router.HandleFunc("/book/", controllers.GetbBook).Methods("GET")            //if someone come to my/book/ o ge all the books that exist send him/her  to getbook function to deal with
// 	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET") // get book with the id that the use send
// 	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
// 	router.HandleFunc("/book/{bookId}", controllers.DeletebBook).Methods("DELETE")

// } //func
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
