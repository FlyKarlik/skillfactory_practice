package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDb(databaseMongoURL string, dataBaseMongoName string) (*mongo.Database, error) {
	clientMongo, err := mongo.NewClient(options.Client().ApplyURI(databaseMongoURL))
	if err != nil {
		return nil, err
	}
	if err = clientMongo.Connect(context.Background()); err != nil {
		return nil, err
	}
	if err = clientMongo.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	return clientMongo.Database(dataBaseMongoName), nil
}
