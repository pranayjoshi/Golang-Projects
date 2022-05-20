package router

import (
	"github.com/gorilla/mux"
	"github.com/pranayjoshi/golang-projects/crudapi/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie/{id}", controller.GetMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", controller.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie", controller.AddMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")

	return r
}
