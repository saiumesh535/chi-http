package database

import (
	mongodb "github.com/globalsign/mgo"
)

var dbConnection *mongodb.Session

// EstablishConnection with MongoDB
func EstablishConnection() {
	url := "mongodb://saiumesh:saiumesh535@ds111410.mlab.com:11410/chi-http"
	session, err := mongodb.Dial(url)
	if err != nil {
		panic("Error in connecting database!!")
	} else {
		dbConnection = session
	}
}
