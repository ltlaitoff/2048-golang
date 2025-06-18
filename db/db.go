package db

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongoDb(url string) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(url))

	if err != nil {
		return nil, err
	}

	slog.Info("MongoClient connected")

	return client, nil
}
