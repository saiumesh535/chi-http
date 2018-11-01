package models


import (
	db "../database"
)

// FindAll mongoDb query
func FindAll(query interface{}, result interface{}) error {
	return db.DbConnection.DB("chi-http").C("users").Find(query).All(result)
}


// FindOne mongoDb query
func FindOne(query interface{}, result interface{}) (err error) {
	return db.DbConnection.DB("chi-http").C("users").Find(query).One(result)
}
