package container

import (
	"fmt"
	rcplgc "github.com/yael-castro/cb-search-engine-api/internal/recipes/business/logic"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input/handler"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output/finder"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output/storage"
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
	recipeCollection := db.Collection("recipes")

	// Secondary adapters
	recipeFinder := finder.NewRecipeFinder(recipeCollection)
	recipeStore := storage.NewRecipeStore(recipeCollection)

	// Ports for primary adapters
	recipeSearcher := rcplgc.NewRecipeSearcher(recipeFinder)
	recipeManager := rcplgc.NewRecipeManager(recipeStore)

	// Primary adapters
	searcher := handler.NewRecipeEngine(recipeSearcher, handler.ErrorHandler())
	creator := handler.NewRecipeCreator(recipeManager, handler.ErrorHandler())

	// Builds HTTP server
	*h = server.New(server.Config{
		RouteMaps: []server.RouteMap{
			handler.RouteMap(creator, searcher),
		},
	})

	return err
}
