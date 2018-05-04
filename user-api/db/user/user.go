package user

import (
	"github.com/Oscar170/i-rol-back/user-api/models"

	"github.com/Oscar170/i-rol-back/utils/db/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	userCollection *mgo.Collection
)

func init() {
	collectionName := "user"
	collecetions, _ := mongo.DB.CollectionNames()
	userCollection = mongo.DB.C(collectionName)

	includes := func(list []string, itemToFind string) bool {
		for _, item := range list {
			if item == itemToFind {
				return true
			}
		}
		return false
	}
	if !includes(collecetions, collectionName) {
		userCollection.EnsureIndex(mgo.Index{
			Key:    []string{"username"},
			Unique: true,
		})

		userCollection.EnsureIndex(mgo.Index{
			Key:    []string{"email"},
			Unique: true,
		})
	}
}

func Store(u *models.User) error {
	return userCollection.Insert(u)
}

func IsValid(u models.User) (int, error) {
	return userCollection.Find(bson.M{
		"$or": []bson.M{
			bson.M{"email": u.Email},
			bson.M{"username": u.Username},
		},
	}).Count()
}

func Restore(u *models.User, c *models.Credentials) error {
	return userCollection.Find(bson.M{
		"username": c.Username,
		"password": c.Password,
	}).One(u)
}
