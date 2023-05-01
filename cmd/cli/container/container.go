package container

import (
	"errors"
	"flag"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/logic"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/infrastructure/input/command"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/infrastructure/output/storage"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/cli"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/connection"
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

	db, err := connection.NewMongoDatabase(mongoDSN, mongoDB)
	if err != nil {
		return errors.New("failed connection: failed connection to external storage")
	}

	recipeCollection := db.Collection("recipes")

	recipeStore := storage.NewRecipeStore(recipeCollection)
	recipeSetGenerator := logic.NewRecipeSetGenerator(recipeStore)
	recipeListGenerator := command.NewRecipeListGenerator(recipeSetGenerator)

	config := cli.Configuration{
		Version:     GitCommit,
		Description: "is a tool for managing the recipes storage and can test the search engine.",
		Commands: map[string]cli.Command{
			"generate": recipeListGenerator.GenerateRecipeList,
		},
		FlagSet: flag.NewFlagSet(os.Args[0], flag.ContinueOnError),
	}

	*cl = cli.New(config)

	return nil
}
