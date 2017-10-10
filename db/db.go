package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	session, err := mgo.Dial("mongo")
	if err != nil {
		panic(err)
	}
	log.Println("Connected to mongo successfully")
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("ignoshi_snippets")
}

// GetDB returns object of the database connection
func GetDB() *mgo.Database {
	return db
}
