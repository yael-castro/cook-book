package container

import (
	"fmt"
	"github.com/rs/cors"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/pkg/connection"
	"github.com/yael-castro/cb-search-engine-api/pkg/server"
	"net/http"
	"os"
)

var GitCommit = ""

type Container interface {
	Inject(any) error
}

func New() Container {
	return container{}
}

type container struct{}

func (c container) Inject(a any) error {
	h, ok := a.(*http.Handler)
	if !ok {
		return fmt.Errorf("type \"%T\" is not supported", a)
	}

	// Establishes connection to MongoDB
	db, err := connection.NewMongoDatabase(os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB"))
	if err != nil {
		return err
	}

	// MongoDB collections
	recipesCollection := db.Collection("recipes")

	// Secondary adapters
	recipeFinder := output.NewRecipesSearcher(recipesCollection)
	recipeCreator := output.NewRecipeCreator(recipesCollection)

	// Ports for primary adapters
	recipeSearcher := business.NewRecipesFinder(recipeFinder)
	recipeAdder := business.NewRecipeAdder(recipeCreator)

	// Primary adapters
	searcher := input.NewRecipesFinder(recipeSearcher, input.ErrorHandler())
	creator := input.NewRecipesCreator(recipeAdder, input.ErrorHandler())

	// Builds HTTP server
	*h = server.New(server.Config{
		RouteMaps: []server.RouteMap{
			input.RouteMap(creator, searcher),
		},
	})

	*h = cors.Default().Handler(*h)
	return err
}
