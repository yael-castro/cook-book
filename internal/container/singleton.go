package container

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"sync"
)

var (
	onceMongo     sync.Once
	mongoDatabase *mongo.Database
)

func MongoDatabase(ctx context.Context) (*mongo.Database, error) {
	var err error

	onceMongo.Do(func() {
		dsn, dbName := os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB")

		var mongoClient *mongo.Client
		mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(dsn))
		if err != nil {
			return
		}

		err = mongoClient.Ping(ctx, readpref.Primary())
		if err != nil {
			return
		}

		database := mongoClient.Database(dbName)

		err = database.Client().Ping(ctx, readpref.Primary())
		if err != nil {
			return
		}

		mongoDatabase = new(mongo.Database)
	})

	return mongoDatabase, nil
}
