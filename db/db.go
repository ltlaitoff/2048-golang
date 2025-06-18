package db

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Database *mongo.Client

func ConnectMongoDb(url string) error {
	client, err := mongo.Connect(options.Client().ApplyURI(url))

	if err != nil {
		return err
	}

	Database = client
	slog.Info("MongoClient connected")

	return nil
}
