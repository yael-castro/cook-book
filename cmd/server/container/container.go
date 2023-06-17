package container

import (
	"fmt"
	"github.com/rs/cors"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/logic"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input/handler"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output/reads"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output/writes"
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
	recipeFinder := reads.NewRecipesSearcher(recipesCollection)
	recipeCreator := writes.NewRecipeCreator(recipesCollection)

	// Ports for primary adapters
	recipeSearcher := logic.NewRecipesFinder(recipeFinder)
	recipeAdder := logic.NewRecipeAdder(recipeCreator)

	// Primary adapters
	searcher := handler.NewRecipesFinder(recipeSearcher, handler.ErrorHandler())
	creator := handler.NewRecipesCreator(recipeAdder, handler.ErrorHandler())

	// Builds HTTP server
	*h = server.New(server.Config{
		RouteMaps: []server.RouteMap{
			handler.RouteMap(creator, searcher),
		},
	})

	*h = cors.Default().Handler(*h)
	return err
}
