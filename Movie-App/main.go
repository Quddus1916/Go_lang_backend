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
	ID       string    `json:"ID"`
	Title    string    `json:"Title"`
	Director *Director `json:"Director"`
}

type Director struct {
	Name string `json:"name"`
}

var movies []Movie

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, items := range movies {

		if items.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, items := range movies {
		if items.ID == params["ID"] {
			json.NewEncoder(w).Encode(items)
			return
		}
	}
}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, items := range movies {

		if items.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}

func main() {
	mux := mux.NewRouter()
	port := "8080"
	mux.HandleFunc("/movies", getmovies).Methods("GET")
	mux.HandleFunc("/movies/{ID}", getmovie).Methods("GET")
	mux.HandleFunc("/movies", createmovie).Methods("POST")
	mux.HandleFunc("/movies/{ID}", updatemovie).Methods("PUT")
	mux.HandleFunc("/movies/{ID}", deletemovie).Methods("DELETE")

	movies = append(movies, Movie{ID: "1", Title: "superman", Director: &Director{Name: "john"}})
	movies = append(movies, Movie{ID: "2", Title: "batman", Director: &Director{Name: "polly"}})

	fmt.Printf("server is running on port %v", port)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
