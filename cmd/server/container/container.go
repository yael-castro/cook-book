package container

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/pkg/connection"
	"github.com/yael-castro/cb-search-engine-api/pkg/server"
	"log"
	"net/http"
	"os"
)

var GitCommit = ""

func Inject(a any) error {
	h, ok := a.(*http.Handler)
	if !ok {
		return fmt.Errorf("type \"%T\" is not supported", a)
	}

	logger := log.Default()

	// Establishes connection to MongoDB
	db, err := connection.NewMongoDatabase(os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB"))
	if err != nil {
		return err
	}

	// MongoDB collections
	recipesCollection := db.Collection("recipes")

	// Secondary adapters
	recipeFinder := output.NewRecipesSearcher(recipesCollection)
	recipeCreator := output.NewRecipeCreator(db, logger)

	// Ports for primary adapters
	recipeSearcher := business.NewRecipesFinder(recipeFinder)
	recipeAdder := business.NewRecipeAdder(recipeCreator)

	// Primary adapters
	searcher := input.NewRecipesFinder(recipeSearcher, input.ErrorHandler())
	creator := input.NewRecipesCreator(recipeAdder, input.ErrorHandler())

	// Builds HTTP server
	*h = server.New(input.RouteMap(creator, searcher))
	return err
}
