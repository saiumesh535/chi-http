package main

import (
	"net/http"

	auth "./auth"
	models "./models"
	renderPkg "github.com/unrolled/render"
)

var render *renderPkg.Render

func init() {
	render = renderPkg.New()
}

// LoadTest MongoDB connections
func LoadTest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var usersData []auth.SignupData
	err := models.FindAll(nil, &usersData)
	if err != nil {
		w.Write([]byte("Something wen't wrong!!"))
	} else {
		render.JSON(w, 200, &usersData)
	}
}
