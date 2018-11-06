package main

import (
	"fmt"
	"net/http"
	db "./database"
	utils "./utils"
	upload "./upload"
	"github.com/go-chi/chi"
	auth "./auth"
)

func init() {
	// establish database connection
	go db.EstablishConnection()
}

// welcome router
func welcome(w http.ResponseWriter, r *http.Request) {
	username := "saiumesh"
	w.Write([]byte(fmt.Sprintf("Article %s has been created", username)))
}

func main() {
	r := chi.NewRouter()
	r.NotFound(utils.Errorhandler)
	r.Get("/", welcome)

	// just for loadtesting purpose, don't use in production
	// running command go run main.go loadTest.go
	// r.Get("/loadtest", LoadTest)

	// all auth routers
	r.Mount("/auth", auth.Handler())

	// all upload files
	r.Mount("/upload", upload.Handler())

	err := http.ListenAndServe(":4321", r)
	if err != nil {
		fmt.Println("Error in creating server ", err)
	}
}
