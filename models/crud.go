package models


import (
	db "../database"
)

// FindAll mongoDb query
func FindAll(query interface{}, result interface{}) error {
	dbConn := db.GetDatabaseConnObject()
	return dbConn.DB("chi-http").C("users").Find(query).All(result)
}
