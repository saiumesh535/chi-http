package models


import (
	db "../database"
)

// FindAll mongoDb query
func FindAll(query interface{}, result interface{}) error {
	return db.DbConnection.DB("chi-http").C("users").Find(query).All(result)
}


// FindOne mongoDb query
func FindOne(query interface{}, result interface{}, collection string) (err error) {
	return db.DbConnection.DB("chi-http").C(collection).Find(query).One(result)
}

// Insert data into database
func Insert(input interface{}, collection string ) (err error) {
	return db.DbConnection.DB("chi-http").C(collection).Insert(input)
}