package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pranayjoshi/golang-projects/emailverifier/router"
)

func main() {
	fmt.Println("Email Checker")
	fmt.Println("Server is Starting at localhost:4000...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening")
}
