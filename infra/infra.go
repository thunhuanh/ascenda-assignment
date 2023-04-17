package infra

import (
	"ascenda-assignment/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB() {
}

// import mongodb driver

var _dbClient *mongo.Client

func GetDBClient() *mongo.Client {
	if _dbClient == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DB_URI))
		if err != nil {
			panic(err)
		}
		_dbClient = client
		return _dbClient

	}
	return _dbClient
}
