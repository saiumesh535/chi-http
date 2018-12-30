package private

import (
	"fmt"
	"net/http"
)

// GetUsers ->  getting All Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Article %s has been created", r.Header.Get("userId"))))
}
