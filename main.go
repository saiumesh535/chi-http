package main

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"

	renderPkg "github.com/unrolled/render"

	jwt "github.com/dgrijalva/jwt-go"

	db "./database"
	utils "./utils"
	"github.com/go-chi/chi"
)

type loginData struct {
	Token string `json:"token"`
}

type jwtCliams struct {
	username string
	jwt.StandardClaims

}

var hmacSampleSecret []byte

var render *renderPkg.Render

func main() {
	hmacSampleSecret = []byte("my_secret_key")
	render = renderPkg.New()
	db.EstablishConnection()
	r := chi.NewRouter()
	r.NotFound(utils.Errorhandler)
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		users := db.ConnectMongoDB()
		utils.JsonHeaders(w)
		json.NewEncoder(w).Encode(users)
	})
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	r.Get("/users", db.GetUsers)
	r.Post("/wtf", func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		username := req.FormValue("username")
		if username == "saiumesh" {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCliams {
				username: username,
				StandardClaims: jwt.StandardClaims {
					ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
				},
			})
			tokenString, err := token.SignedString(hmacSampleSecret)
			if err == nil {
				loginReponse := loginData{
					Token: tokenString,
				}
				render.JSON(w, 200, loginReponse)
			}
		} else {
			w.Write([]byte("something wen't wrong"))
		}
	})
	err := http.ListenAndServe(":4321", r)
	if err != nil {
		fmt.Println("Error in creating server ", err)
	}
}
