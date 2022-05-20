package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Director Director `json:"director"`
	Actor    *Actor   `json:"actor"`
}

type Director struct {
	Name string `json:"name"`
	Exp  int    `json:"exp"`
}

type Actor struct {
	Name  string `json:"name"`
	Oscar *bool  `json:"oscar"`
}

var movies []Movie

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("no data found")
	}

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	for _, m := range movies {
		if movie.Name == m.Name {
			json.NewEncoder(w).Encode("Movie Already exists")
		}
	}

	rand.Seed(time.Now().Unix())
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode("Movie Successfully Deleted!")
			break
		}
	}

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode("Movie Successfully Updated!")
			break
		}
	}

}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			movies = nil
			json.NewEncoder(w).Encode("Movies Database Successfully Deleted!")
			break
		}
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode("No records!")
}
