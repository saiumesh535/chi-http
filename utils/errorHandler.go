package utils

import (
	"net/http"

	renderPkg "github.com/unrolled/render"
)

//ErrorMessage for 404
type errorMessage struct {
	Status  bool
	Message string
}

var render *renderPkg.Render

func init() {
	render = renderPkg.New()
}

// Errorhandler for 404
func Errorhandler(w http.ResponseWriter, r *http.Request) {
	errorMessage := errorMessage {
		Status: false,
		Message: "Doomed!!",
	}
	render.JSON(w, 404, errorMessage)
}
