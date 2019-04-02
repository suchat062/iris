package config

import (
	mgo "gopkg.in/mgo.v2"
	"os"
)

func GetMongoDB() (*mgo.Database, error) {
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}