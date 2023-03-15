package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/Json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/Json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/Json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/Json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/Json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func main() {

	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn: "Het2423", Title: "Jack Saparow",
		Director: &Director{
			Id:        "1",
			FirstName: "Palazo",
			LastName:  "Jiggy",
		}})
	movies = append(movies, Movie{Id: "2", Isbn: "Het24423", Title: "Saparow",
		Director: &Director{
			Id:        "1",
			FirstName: "Nice",
			LastName:  "Jiggy",
		}})

	r.HandleFunc("/movies/", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}/", getMovie).Methods("GET")
	r.HandleFunc("/movie/", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}/update/", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}/delete/", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at port 7070 \n")
	log.Fatal(http.ListenAndServe(":7070", r))

}
