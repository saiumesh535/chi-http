package auth

import (
	"github.com/globalsign/mgo/bson"
	"encoding/json"
	"net/http"

	renderPkg "github.com/unrolled/render"

	models "../models"
	utils "../utils"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Loginresponse structure
type Loginresponse struct {
	Token string `json:"token"`
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

	loginData := bson.M{
		"username": r.FormValue("username"),
	}
	err := models.FindOne(&loginData, &result, "users")
	if err != nil && err.Error() != "not found" {
		w.Write([]byte("Something wen't wrong"))
	} else {
		if len(result.Username) == 0 {
			render.JSON(w, 403, map[string]interface{}{
				"status":  false,
				"message": "yo!! check username and password",
			})
		} else {
			isVliadPassword := utils.CheckPasswordHash(r.FormValue("password"), result.Password)
			if isVliadPassword {
				out, _ := json.Marshal(result)
				token, err := utils.CreateJwtToken(string(out))
				if err != nil {
					w.Write([]byte("Token wen't wrong"))
				} else {
					response := Loginresponse{
						Token: token,
					}
					render.JSON(w, 200, &response)
				}
			} else {
				render.JSON(w, 403, map[string]interface{}{
					"status":  false,
					"message": "yo!! check username and password",
				})
			}

		}

	}
}
