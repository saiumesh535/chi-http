package database

import (
	mongodb "github.com/globalsign/mgo"
)

var dbConnection *mongodb.Session

// EstablishConnection with MongoDB
func EstablishConnection() {
	url := "mongodb://localhost:27017/chi-http"
	session, err := mongodb.Dial(url)
	if err != nil {
		panic("Error in connecting database!!")
	} else {
		dbConnection = session
	}
}

// GetDatabaseConnObject returns database connection
func GetDatabaseConnObject() *mongodb.Session {
	return dbConnection;
}
