package connection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoDatabase establish connection with a MongoDB database using the Configuration passed as parameter
func NewMongoDatabase(dsn, database string) (mongoDatabase *mongo.Database, err error) {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		return
	}

	err = mongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return
	}

	mongoDatabase = mongoClient.Database(database)
	err = mongoDatabase.Client().Ping(context.TODO(), readpref.Primary())
	return
}
