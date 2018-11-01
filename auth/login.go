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
	var result login
	r.ParseForm()
	loginData := login{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	err := models.FindOne(&loginData, &result)
	if  err != nil && err.Error() != "not found" {
		w.Write([]byte("Something wen't wrong"))
	} else {
		if len(result.Password) == 0 {
			render.JSON(w, 403, map[string]interface{}{
				"status":  false,
				"message": "yo!! check username and password",
			})
		} else {
			render.JSON(w, 200, result)
		}
	}
}
