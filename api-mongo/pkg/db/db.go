package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetDB() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	database := client.Database("joewalker")

	return database, nil
}