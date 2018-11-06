package database

import (
	mongodb "github.com/globalsign/mgo"
)

// DbConnection database connection
var DbConnection *mongodb.Session

// EstablishConnection with MongoDB
func EstablishConnection() {
	url := "mongodb://localhost:27017/chi-http"
	session, err := mongodb.Dial(url)
	if err != nil {
		panic("Error in connecting database!!")
	} else {
		DbConnection = session
	}
}

// GetDatabaseConnObject returns database connection
func GetDatabaseConnObject() *mongodb.Session {
	return DbConnection;
}
