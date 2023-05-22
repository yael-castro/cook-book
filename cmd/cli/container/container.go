package container

import (
	"errors"
	"flag"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/logic"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input/command"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output/generator"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output/storage"
	"github.com/yael-castro/cb-search-engine-api/pkg/cli"
	"github.com/yael-castro/cb-search-engine-api/pkg/connection"
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

	recipesGenerator := generator.NewRecipesGenerator()
	recipeStore := storage.NewRecipeStore(recipeCollection)
	recipeSetGenerator := logic.NewRecipeSetGenerator(recipesGenerator, recipeStore)
	recipeListGenerator := command.NewRecipeListGenerator(recipeSetGenerator)

	config := cli.Configuration{
		Version:     GitCommit,
		Description: "is a tool for managing the recipes storage and can test the search engine.",
		Commanders: map[string]cli.Commander{
			"generate": recipeListGenerator,
		},
		FlagSet: flag.NewFlagSet(os.Args[0], flag.ContinueOnError),
	}

	*cl = cli.New(config)

	return nil
}
