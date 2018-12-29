package auth

import (
	"fmt"
	"net/http"
	bson "github.com/globalsign/mgo/bson"
	models "../models"
	utils "../utils"
)

// SignupData struct
type SignupData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Signup usres
func Signup(w http.ResponseWriter, r *http.Request) {
	var inputData SignupData
	defer r.Body.Close()
	r.ParseForm()
	// check if email present or not
	query := bson.M{
		"$or": []interface{}{
			bson.M{"email": r.FormValue("email")},
			bson.M{"username": r.FormValue("username")},
		},
	}
	err := models.FindOne(query, &inputData, "users")

	if err != nil && err.Error() != "not found" {
		w.Write([]byte("Something wen't wrong"))
	} else if len(inputData.Email) != 0 {
		render.JSON(w, 403, map[string]interface{}{
			"status":  false,
			"message": "Email or Username already exists!!",
		})
	} else {
		hash, hashErr := utils.HashPassword(r.FormValue("password"))
		if hashErr != nil {
			w.Write([]byte("Hash Error"))
		} else {
			inputData = SignupData{
				Username: r.FormValue("username"),
				Password: hash,
				Email:    r.FormValue("email"),
			}
			err := models.Insert(&inputData, "users")

			if err != nil {
				w.Write([]byte("Something wen't wrong!!"))
			} else {
				fmt.Fprintf(w, "Welcome to golang %s", inputData.Username)
			}
		}
	}
}
