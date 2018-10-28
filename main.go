package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "./database"
	utils "./utils"
	"github.com/go-chi/chi"
)

func main() {
	db.EstablishConnection()
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users := db.ConnectMongoDB()
		utils.JsonHeaders(w)
		json.NewEncoder(w).Encode(users)
	})
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	r.Get("/users", db.GetUsers)
	err := http.ListenAndServe(":4321", r)
	if err != nil {
		fmt.Println("Error in creating server ", err)
	}
}
