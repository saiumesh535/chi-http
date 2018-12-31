package private

import (
	"fmt"
	"net/http"

	model "../models"
	renderPkg "github.com/unrolled/render"
)

type UsersData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	render = renderPkg.New()
}

// GetUsers ->  getting All Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []UsersData
	err := model.FindAll(nil, &users)
	if err != nil {
		fmt.Println("err", err)
		w.Write([]byte("Something wen't wrong!!"))
	} else {
		render.JSON(w, 200, &users)
	}
}
