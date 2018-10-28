package database

import (
	mongodb "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type users struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

func ConnectMongoDB() []users {
	url := "33"
	session, err := mongodb.Dial(url)
	if err != nil {
		return nil
	} else {
		var result []users
		collection := session.DB("chi-http").C("users")
		iter := collection.Find(bson.M{}).Limit(100).Iter()
		iter.All(&result)
		return result
	}
}
