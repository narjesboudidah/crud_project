package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string `json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	Director string `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie


func getMovies(w http.Responsewriter , r *http.Request) 
{
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie( w http.Responsewriter , r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies {
		if item.ID==params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			break 
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.Responsewriter, r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r) 
	for _,item:=range movies {
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}
func createMovie(w http.Responsewriter, r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_=json.NewEncoder(r,Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(100000000))
	movies=append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.Responsewriter, r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies {
	if item.ID==params["id"]{
		movies=append(movies[:index],movies[index+1:]...)
		var movie Movie 
		_=json.NewEncoder(r,Body).Decode(&movie)
		movie.ID=params["id"]
		movies=append(movies,movie)
		json.NewEncoder(w).Encode(movie)


	}
	}
}
func main() {
	r := mux.NewRouter()

	movies=append(movies,Movie[ID:"1",Isbn:"438227",Title:"Movie One",
	Director:&Director{Firstname:"David",Lastname:"Amelson"}])
	movies=append(movies,Movie[ID:"2",Isbn:"425896",Title:"Second Movie",Director:&Director{Firstname:"Alice",Lastname:"Tals"}])

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Start Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000"r))
}
