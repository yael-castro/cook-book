package container

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

// GitCommit is the git hash from the binary was built.
var GitCommit = ""

func inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case *mongo.Database:
		return injectMongoDatabase(ctx, a)
	}

	return fmt.Errorf("type \"%T\" is not supported", a)
}

func injectMongoDatabase(ctx context.Context, database *mongo.Database) (err error) {
	dsn, dbName := os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB")

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	*database = *mongoClient.Database(dbName)

	err = database.Client().Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	return
}
