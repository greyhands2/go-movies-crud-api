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
	ID       string    `json:"id"`
	Isbn     string    `json:isbn`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(movies)

}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "applicaiton/json")

	params := mux.Vars(req)

	for index, val := range movies {
		if val.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break

		}

	}
	defer json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, e := range movies {
		if e.ID == params["id"] {
			json.NewEncoder(res).Encode(e)
			return
		}
	}

}

func createMovie(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000000))

	movies = append(movies, movie)

	json.NewEncoder(res).Encode(movie)

}

func updateMovie(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	var movie Movie
	for i, e := range movies {
		if e.ID == params["id"] {
			//delete movies first
			movies = append(movies[:i], movies[i+1:]...)

			_ = json.NewDecoder(req.Body).Decode(&movie)

			movie.ID = params["id"]

			movies = append(movies, movie)
			break

		}
	}

	defer json.NewEncoder(res).Encode(movie)
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "3435343", Title: "John Wick", Director: &Director{Firstname: "Osas", Lastname: "Tony"}})

	movies = append(movies, Movie{ID: "2", Isbn: "3435344", Title: "Harry Potter", Director: &Director{Firstname: "John", Lastname: "Buskle"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")

	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	router.HandleFunc("/movies", createMovie).Methods("POST")

	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting serverat 4800")

	log.Fatal(http.ListenAndServe(":4800", router))

}
