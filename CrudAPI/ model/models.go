package model

import "fmt"

type Movie struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Actor    string `json:"actor"`
}

type Director struct {
	Name string `json:"name"`
	Exp  int    `json:"exp"`
}

type Actor struct {
	Name  string `json:"name"`
	Oscar bool   `json:"oscar"`
}

func main() {
	fmt.Println("Models")
}
