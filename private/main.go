package private

import (
	"encoding/json"
	"net/http"

	utils "../utils"
	"github.com/go-chi/chi"
	"github.com/mitchellh/mapstructure"
	renderPkg "github.com/unrolled/render"
)

// MongoDBID of user
type MongoDBID struct {
	ID string `json:"id"`
}

// JwtData of user
type JwtData struct {
	Data string
	Iss  string `json:"iss"`
}

var render *renderPkg.Render

func init() {
	render = renderPkg.New()
}

func checkJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			render.JSON(w, http.StatusForbidden, map[string]interface{}{
				"status":  false,
				"message": "token is invalid",
			})
		} else {
			decode, _ := utils.DecodeJwtToken(token)
			var jwt JwtData
			mapstructure.Decode(decode, &jwt)
			var user MongoDBID
			json.Unmarshal([]byte(jwt.Data), &user)
			r.Header.Set("userId", user.ID)
			next.ServeHTTP(w, r)
		}
	})
}

// Handler for all auth purposes
func Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(checkJwt)
	r.Get("/getAllUsers", GetUsers)
	return r
}
