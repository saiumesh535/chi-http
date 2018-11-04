package auth


import (
	"fmt"
	"net/http"

	models "../models"
)

// SignupData struct
type SignupData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

// Signup usres
func Signup(w http.ResponseWriter, r *http.Request) {
	var inputData SignupData
	defer r.Body.Close();
	r.ParseForm();
	// check if email present or not
	err := models.FindOne(map[string]interface{}{
		"email": r.FormValue("email"),
	}, &inputData, "users")

	if  err != nil && err.Error() != "not found" {
		w.Write([]byte("Something wen't wrong"))
	} else if len(inputData.Email) != 0 {
		render.JSON(w, 403, map[string]interface{}{
			"status": false,
			"message": "Email already exists!!",
		})
	} else {
		inputData = SignupData {
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email: r.FormValue("email"),
		}
		err := models.Insert(&inputData, "users")

		if err != nil {
			w.Write([]byte("Something wen't wrong!!"))
		} else {
			fmt.Fprintf(w, "Welcome to golang %s", inputData.Username)
		}

	}
}