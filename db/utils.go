package db

import (
	"gopkg.in/mgo.v2"
	"os"
)

var session *mgo.Session

func connect() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(getConnectionString())
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
	}
	return session
}

const localConnection = "mongodb://localhost/cooking"

func getConnectionString() string {
	connectString := os.Getenv("DB_CONNECTION")
	if connectString == "" {
		return localConnection
	}
	return connectString
}
