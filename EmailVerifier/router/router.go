package router

import (
	"github.com/gorilla/mux"
	"github.com/pranayjoshi/golang-projects/emailverifier/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/email", controller.CheckDomain)
	return r
}
