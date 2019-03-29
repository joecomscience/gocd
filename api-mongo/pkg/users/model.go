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

func (u *User) GetUsers(conn *mongo.Database) ([]User, error) {
	var users []User

	collection := conn.Collection("users")
	data, err := collection.Find(context.Background(), bson.D{}, nil)
	if err != nil {
		return nil, err
	}
	defer data.Close(context.Background())

	for data.Next(context.Background()) {
		var tempUser User
		if err := data.Decode(&tempUser); err != nil {
			log.Fatal(err)
		}

		users = append(users, tempUser)
	}

	if err := data.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) CreateUser(conn *mongo.Database, user User) (*mongo.InsertOneResult, error) {
	collection := conn.Collection("users")
	user.ID = primitive.NewObjectID()

	data, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return data, nil
}
