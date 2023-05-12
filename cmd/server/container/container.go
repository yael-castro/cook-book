package container

import (
	"fmt"
	rcplgc "github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/logic"
	rcphlr "github.com/yael-castro/cb-search-engine-api/internal/core/recipes/infrastructure/input/handler"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/infrastructure/output/storage"
	fndlgc "github.com/yael-castro/cb-search-engine-api/internal/core/searches/business/logic"
	fndhlr "github.com/yael-castro/cb-search-engine-api/internal/core/searches/infrastructure/input/handler"
	"github.com/yael-castro/cb-search-engine-api/internal/core/searches/infrastructure/output/finder"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/connection"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/server"
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
	recipeSearcher := fndlgc.NewRecipeSearcher(recipeFinder)
	recipeManager := rcplgc.NewRecipeManager(recipeStore)

	// Primary adapters
	recipeEngine := fndhlr.NewRecipeEngine(recipeSearcher, fndhlr.ErrorHandler())
	recipeCreator := rcphlr.NewRecipeCreator(recipeManager, rcphlr.ErrorHandler())

	// Builds HTTP server
	*h = server.New(server.Config{
		RouteMaps: []server.RouteMap{
			fndhlr.RouteMap(recipeEngine),
			rcphlr.RouteMap(recipeCreator),
		},
	})

	return err
}
