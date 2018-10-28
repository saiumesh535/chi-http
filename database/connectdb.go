package database

import (
	"net/http"

	mongodb "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type users struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

var dbConnection *mongodb.Session

func EstablishConnection() {
	url := "mongodb://saiumesh:saiumesh535@ds111410.mlab.com:11410/chi-http"
	session, err := mongodb.Dial(url)
	if err != nil {
		panic("Error in connecting database!!")
	} else {
		dbConnection = session
	}
}

func ConnectMongoDB() []users {
	var result []users
	collection := dbConnection.DB("chi-http").C("users")
	iter := collection.Find(bson.M{})
	iter.All(&result)
	return result
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
