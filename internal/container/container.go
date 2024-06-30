package container

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
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
	db, err := MongoDatabase(ctx)
	if err != nil {
		return err
	}

	*database = *db
	return
}
