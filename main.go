package main

import (
	"fmt"
	"net/http"
	db "./database"
	utils "./utils"
	"github.com/go-chi/chi"
)

func init() {
	// establish database connection
	go db.EstablishConnection()
}

// welcome router
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Golang develpment!!"))
}

func main() {
	r := chi.NewRouter()
	r.NotFound(utils.Errorhandler)
	r.Get("/", welcome)
	err := http.ListenAndServe(":4321", r)
	if err != nil {
		fmt.Println("Error in creating server ", err)
	}
}
