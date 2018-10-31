package auth

import (
	"net/http"

	renderPkg "github.com/unrolled/render"

	models "../models"
)

type login struct {
	Username string
	Password string
}

var render *renderPkg.Render

func init() {
	render = renderPkg.New()
}

// Login users
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var result []login
	r.ParseForm()
	loginData := login{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	err := models.FindAll(&loginData, &result)
	if err != nil {
		w.Write([]byte("Something wen't wrong"))
	} else {
		if result == nil {
			render.JSON(w, 403, map[string]interface{}{
				"status":  false,
				"message": "yo!! check username and password",
			})
		} else {
			render.JSON(w, 200, result)
		}
	}
}
