package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	db "./database"
	download "./download"
	utils "./utils"
	public "./static"
	upload "./upload"
	"github.com/go-chi/chi"
	auth "./auth"
	cors "github.com/go-chi/cors"
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

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // you can add routes here www.example.com
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	r.NotFound(utils.Errorhandler)
	r.Get("/", welcome)

	// sending static files
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	// now visit http://localhost:4321/public/
	// it will serve index.html
	// for http://localhost:4321/public/someother.html
	// it will serve someother.html file
	public.FileServer(r, "/public", http.Dir(filesDir))

	// just for loadtesting purpose, don't use in production
	// running command go run main.go loadTest.go
	// r.Get("/loadtest", LoadTest)

	// all auth routers
	r.Mount("/auth", auth.Handler())

	// all upload files
	r.Mount("/upload", upload.Handler())

	r.Get("/downloadFile", download.SmallFiles)

	err := http.ListenAndServe(":4321", r)
	if err != nil {
		fmt.Println("Error in creating server ", err)
	}
}
