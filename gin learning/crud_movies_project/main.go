package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand" // create a id to each movie
	"net/http"  // create a server in golang
	"strconv"

	"github.com/gorilla/mux"
)

//struct and slices will replace database

// here struct of type movie
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json :"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// implement getMovies function
func getMovies(w http.ResponseWriter, r *http.Request) { //passing a pointer of the request that you will sent from postman to this function and w is the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) //
}

// delete funtion
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] { //if item.id == the id sent from the request
			//	this movie eon't exist as you put undex+1 in its place
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies) //return all movies after delete specific one
}

//get movie function

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] { //if item.id == the id sent from the request
			//	this movie eon't exist as you put undex+1 in its place
			json.NewEncoder(w).Encode(item) //retun the movie
			return
		}

	}

}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie                                  //movie is the movie that we will create to send it to postman
	_ = json.NewDecoder(r.Body).Decode(&movie)       //use this to decode the json body that you will send to postman
	movie.ID = strconv.Itoa(rand.Intn(100000000000)) //create new id for the new movies and chooose rand num for id  then convert it to string
	movies = append(movies, movie)                   //append new movie to movies slice
	json.NewEncoder(w).Encode(movie)                 //return this movies that we just created
} //createMovie

// delete first delete that movie which has the id that we will sent then add the id  to the new movie that we will send
// we will delete the name of the old movie and replace it with the new film ,so the new film will have the same id as the old one
func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	// loop over the movies
	for index, item := range movies {
		if item.ID == params["id"] {
			//delete the movie of with the id that we sent
			movies = append(movies[:index], movies[index+1:]...) //to delete movie

			//add new movie which will be the movie that we will send with posman

			var movie Movie                            //movie is the movie that we will create to send it to postman
			_ = json.NewDecoder(r.Body).Decode(&movie) //use this to decode the json body that you will send to postman
			movie.ID = params["id"]                    //create new id for the new movies and chooose rand num for id  then convert it to string
			movies = append(movies, movie)             //append new movie to movies slice
			json.NewEncoder(w).Encode(movie)
			return

		} //if
	} //for

} //update movie
func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director{Firstname: "Jhon", Lastname: "Doe"}})   //append somemovies to test functionality
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie two", Director: &Director{Firstname: "Steve", Lastname: "Smath"}}) //append somemovies to test functionality

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("Get") // movies routes,if I hit /movies route will take you to getMovies function,and the used method is get
	r.HandleFunc("/movies/{id}", getMovie).Methods("Get")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("sRATING SERVER AT PORT 8000")
	log.Fatal(http.ListenAndServe("8000", r)) //log out if the server doesn't start

} //main
