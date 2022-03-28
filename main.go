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

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Status struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type Movie struct {
	Status   *Status   `json:"status"`
	ID       int       `json:"id"`
	ISBN     int       `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Res struct {
	Status *Status `json:"status"`
	Movies []Movie `json:"movies"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var status Status
	status.Success = true

	// Encode movies list to JSON
	json.NewEncoder(w).Encode(&Res{
		Status: &status,
		Movies: movies,
	})
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, m := range movies {
		if strconv.Itoa(m.ID) == params["id"] {
			json.NewEncoder(w).Encode(m)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for i, m := range movies {
		if strconv.Itoa(m.ID) == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = rand.Intn(1000000)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for i, m := range movies {

	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: 1, ISBN: 492338, Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Smith"}})
	movies = append(movies, Movie{ID: 2, ISBN: 457422, Title: "Movie Two", Director: &Director{FirstName: "Mad", LastName: "Max"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", addMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Go CRUD API Server on Port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
