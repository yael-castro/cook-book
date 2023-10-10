package container

import (
	"flag"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/pkg/cli"
	"github.com/yael-castro/cb-search-engine-api/pkg/connection"
	"go.mongodb.org/mongo-driver/mongo"
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

func Inject(a any) error {
	cl, ok := a.(*cli.CLI)
	if !ok {
		return fmt.Errorf("type \"%T\" is not supported", a)
	}

	if mongoDSN == "" {
		mongoDSN = defaultMongoDSN
	}

	if mongoDB == "" {
		mongoDB = defaultMongoDB
	}

	// Dependencies
	logger := log.Default()

	databaseFunc := func() (*mongo.Database, error) {
		return connection.NewMongoDatabase(mongoDSN, mongoDB)
	}

	// Driven adapters
	recipesMaker := output.NewRecipesMaker()
	recipeCreator := output.NewRecipeCreatorFunc(databaseFunc, logger)

	// Ports for driving adapters
	recipeSetGenerator := business.NewRecipeGenerator(recipesMaker, recipeCreator)

	// Driving adapters
	recipeListGenerator := input.NewRecipeListGenerator(recipeSetGenerator)

	config := cli.Configuration{
		Version:     gitCommit,
		Description: "is a tool for managing the recipes writes and can test the search engine.",
		Commanders: map[string]cli.Commander{
			"generate": recipeListGenerator,
		},
		FlagSet: flag.NewFlagSet(os.Args[0], flag.ContinueOnError),
	}

	*cl = cli.New(config)

	return nil
}
