package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pranayjoshi/golang-projects/crudapi/router"
)

func main() {
	fmt.Println("Crud API- Movies")
	fmt.Println("Server is Starting..")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at localhost:400")
}
