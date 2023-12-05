package container

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/pkg/command"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	gitCommit = ""
	mongoDB   = ""
	mongoDSN  = ""
)

const (
	defaultMongoDB  = "test"
	defaultMongoDSN = "mongodb://localhost:27017"
)

func Inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case *command.Command:
		return injectCLI(ctx, a)
	case *mongo.Database:
		return injectMongoDatabase(ctx, a)
	case *func() (*mongo.Database, error):
		return injectDatabaseFunc(ctx, a)
	}

	return fmt.Errorf("type \"%T\" is not supported", a)
}

func injectCLI(ctx context.Context, cl *command.Command) (err error) {
	logger := log.Default()

	var databaseFunc func() (*mongo.Database, error)

	if err := Inject(ctx, &databaseFunc); err != nil {
		return err
	}

	// Driven adapters
	recipesMaker := output.NewRecipesWriter()
	recipeCreator := output.NewRecipeSaverFunc(databaseFunc, logger)

	// Ports for driving adapters
	recipeSetGenerator := business.NewRecipeGenerator(recipesMaker, recipeCreator)

	// Driving adapters
	recipeListGenerator := input.CommandGenerateRecipes(recipeSetGenerator)

	config := command.Configuration{
		Version:     gitCommit,
		Description: "tool for managing the recipes writes and can test the search engine.",
		Commands: []command.Command{
			recipeListGenerator,
		},
	}

	*cl = command.New(config)
	return
}

func injectDatabaseFunc(ctx context.Context, databaseFunc *func() (*mongo.Database, error)) error {
	*databaseFunc = func() (*mongo.Database, error) {
		var mongoDB mongo.Database

		if err := Inject(ctx, &mongoDB); err != nil {
			return nil, err
		}

		return &mongoDB, nil
	}
	return nil
}

func injectMongoDatabase(ctx context.Context, database *mongo.Database) (err error) {
	if mongoDSN == "" {
		mongoDSN = defaultMongoDSN
	}

	if mongoDB == "" {
		mongoDB = defaultMongoDB
	}

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDSN))
	if err != nil {
		return
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	*database = *mongoClient.Database(mongoDB)
	return database.Client().Ping(ctx, readpref.Primary())
}
