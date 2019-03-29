package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Age       int64              `json:"age" bson:"age"`
	Addr      string             `json:"addr" bson:"addr"`
}

func (user *User) GetUsers(database *mongo.Database) ([]User, error) {
	var users []User

	collection := database.Collection("users")
	u, err := collection.Find(context.TODO(), bson.M{}, nil)
	if err != nil {
		return nil, err
	}

	for u.Next(context.TODO()) {
		var tempUser User
		if err := u.Decode(&tempUser); err != nil {
			log.Fatal(err)
		}

		users = append(users, tempUser)
	}

	if err := u.Err(); err != nil {
		return nil, err
	}

	if err := u.Close(context.TODO()); err != nil {
		return nil, err
	}

	return users, nil
}
