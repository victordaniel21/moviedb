package main

import (
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Year     string    `json:"year"`
	Genre    string    `json:"genre"`
	Actor    *Actor    `json:"director"`
	Director *Director `json:"actor"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Actor struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var mov []Movie

func main() {
	r := mux.NewRouter() //r for router

	mov = append(mov, Movie{ID: 1, Title: "Project", Year: "2019", Genre: "drama", Director: &Director{FirstName: "Adam", LastName: "First"}, Actor: &Actor{FirstName: "Albert", LastName: "Big"}})
	mov = append(mov, Movie{ID: 2, Title: "Project Two", Year: "2020", Genre: "drama", Director: &Director{FirstName: "Adam", LastName: "First"}, Actor: &Actor{FirstName: "Albert", LastName: "Big"}})
	mov = append(mov, Movie{ID: 2, Title: "Another Project", Year: "2021", Genre: "drama", Director: &Director{FirstName: "Adam", LastName: "First"}, Actor: &Actor{FirstName: "Albert", LastName: "Big"}})

	r.HandleFunc("/movies", pullMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", pullMovie).Methods("GET")
	r.HandleFunc("/movie", pushMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", chaneMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", eraseMovies).Methods("DELETE")
}
