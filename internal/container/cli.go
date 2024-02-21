//go:build cli

package container

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yael-castro/cook-book/internal/app/recipes/business"
	"github.com/yael-castro/cook-book/internal/app/recipes/infrastructure/input/command"
	"github.com/yael-castro/cook-book/internal/app/recipes/infrastructure/output"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
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
	case *cobra.Command:
		return injectCommand(ctx, a)
	case *mongo.Database:
		return injectMongoDatabase(ctx, a)
	case *func() (*mongo.Database, error):
		return injectDatabaseFunc(ctx, a)
	}

	return fmt.Errorf("type \"%T\" is not supported", a)
}

func injectCommand(ctx context.Context, cmd *cobra.Command) (err error) {
	logger := log.Default()

	var databaseFunc func() (*mongo.Database, error)

	if err := Inject(ctx, &databaseFunc); err != nil {
		return err
	}

	// Driven adapters
	recipesWriter := output.NewRecipesWriter()
	recipesSaver := output.NewRecipeSaverFunc(databaseFunc, logger)

	// Ports for driving adapters
	recipesGenerator := business.NewRecipeGenerator(recipesWriter, recipesSaver)

	// Setting command
	*cmd = cobra.Command{
		Use:     os.Args[0],
		Short:   "Tool for managing the recipes operations.",
		Version: gitCommit,
	}

	cmd.CompletionOptions.DisableDefaultCmd = true

	// Setting drive adapters (sub-commands)
	cmd.AddCommand(command.GenerateRecipes(recipesGenerator))

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
