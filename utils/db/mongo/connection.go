package mongo

import (
	"os"

	"gopkg.in/mgo.v2"
)

// DB pointer to the connection of the mongoDb.
var DB *mgo.Database

// Open connecetio to the mongoDb using env var (docker)
func init() {
	mongoEndpoint := os.Getenv("MONGO_ENPOINT")
	mongoDb := os.Getenv("MONGO_DB")

	session, err := mgo.Dial(mongoEndpoint)
	if err != nil {
		panic(err)
	}

	DB = session.DB(mongoDb)
}
