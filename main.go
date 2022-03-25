package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       int       `json:"id"`
	ISBN     int       `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: 1, ISBN: 492338, Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Smith"}})
	movies = append(movies, Movie{ID: 2, ISBN: 457422, Title: "Movie Two", Director: &Director{FirstName: "Mad", LastName: "Max"}})
	r.HandleFunc("/movies", getMovies).Method("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")

	fmt.Printf("Starting Go CRUD API Server on Port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
