package models

import "gopkg.in/mgo.v2/bson"

// User define the user model in the app.
type User struct {
	ID       bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	Name     string          `json:"name" bson:"name"`
	Email    string          `json:"email" bson:"email"`
	Password string          `json:"password" bson:"password"`
	Friends  []bson.ObjectId `json:"friends" bson:"friends"`
	Rooms    []bson.ObjectId `json:"rooms" bson:"rooms"`
}
