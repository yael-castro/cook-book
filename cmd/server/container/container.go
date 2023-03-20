package container

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/application/handler"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/logic"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/infrastructure/finder"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/connection"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server/health"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	s, ok := a.(*server.Server)
	if !ok {
		return fmt.Errorf("type \"%T\" is not supported", a)
	}

	// Establishes connection to MongoDB
	db, err := connection.NewMongoDatabase(os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB"))
	if err != nil {
		return err
	}

	recipeCollection := db.Collection("recipes")

	// Secondary adapters
	recipeFinder := finder.NewRecipeFinder(recipeCollection)

	// Ports for primary adapters
	recipeSearcher := logic.NewRecipeSearcher(recipeFinder)

	// Primary adapters
	recipeEngine := handler.NewRecipeEngine(recipeSearcher, handler.NewErrorHandler())

	// Server settings
	config := server.Configuration{
		RouteMap: map[string]map[string]http.Handler{
			"/v1/recipes": {
				http.MethodGet: recipeEngine,
			},
			"/v1/health/checks": {
				http.MethodGet: http.HandlerFunc(health.NewChecker(health.Config{
					Version: GitCommit,
					PingMap: map[string]health.Ping{
						"mongodb": func(ctx context.Context) error {
							return db.Client().Ping(ctx, readpref.Primary())
						},
					},
				}).Check),
			},
		},
	}

	// Builds HTTP server
	*s = server.New(config)

	return err
}
