package utils

import (
	"net/http"
)

func Errorhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Hey!! looks like dead end"))
}
