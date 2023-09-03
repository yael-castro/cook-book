package container

import (
	"errors"
	"flag"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/pkg/cli"
	"github.com/yael-castro/cb-search-engine-api/pkg/connection"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var (
	GitCommit = ""
	mongoDB   = ""
	mongoDSN  = ""
)

const (
	defaultMongoDB  = "test"
	defaultMongoDSN = "mongodb://localhost:27017"
)

type Container interface {
	Inject(any) error
}

func New() Container {
	return container{}
}

type container struct{}

func (c container) Inject(a any) error {
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

	recipeCollectionFunc := func() (*mongo.Collection, error) {
		db, err := connection.NewMongoDatabase(mongoDSN, mongoDB)
		if err != nil {
			return nil, errors.New("failed connection: failed connection to external writes")
		}

		return db.Collection("recipes"), nil
	}

	recipesMaker := output.NewRecipesMaker()
	recipeCreator := output.NewRecipeCreatorForCLI(recipeCollectionFunc)
	recipeSetGenerator := business.NewRecipeGenerator(recipesMaker, recipeCreator)
	recipeListGenerator := input.NewRecipeListGenerator(recipeSetGenerator)

	config := cli.Configuration{
		Version:     GitCommit,
		Description: "is a tool for managing the recipes writes and can test the search engine.",
		Commanders: map[string]cli.Commander{
			"generate": recipeListGenerator,
		},
		FlagSet: flag.NewFlagSet(os.Args[0], flag.ContinueOnError),
	}

	*cl = cli.New(config)

	return nil
}
