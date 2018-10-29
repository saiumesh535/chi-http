package utils

import (
	"net/http"

	renderPkg "github.com/unrolled/render"
)

type ErrorMessage struct {
	status  bool
	message string
}

var render *renderPkg.Render

func init() {
	render = renderPkg.New() // pass options if you want
}

func Errorhandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  false,
		"message": "Doomed!!",
	}
	render.JSON(w, 404, response)
}
