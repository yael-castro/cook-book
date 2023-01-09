package container

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/application/handler"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/logic"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/infrastructure/finder"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/connection"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server"
	"os"
)

type Container interface {
	Inject(any) error
}

func New() Container {
	return container{}
}

type container struct{}

func (c container) Inject(a any) error {
	switch t := a.(type) {
	case *server.Server:
		return c.injectToServer(t)
	}

	return fmt.Errorf("type \"%T\" is not supported", a)
}

func (c container) injectToServer(s *server.Server) error {
	db, err := connection.NewMongoDatabase(os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB"))
	if err != nil {
		return err
	}

	recipeCollection := db.Collection("recipes")

	recipeFinder := finder.NewRecipeFinder(recipeCollection)
	recipeSearcher := logic.NewRecipeSearcher(recipeFinder)
	recipeProvider := handler.NewRecipeProvider(recipeSearcher)

	config := server.Configuration{
		SearchEngine: recipeProvider.ProvideRecipe,
		ErrorHandler: server.DefaultErrorHandler,
	}

	*s = server.New(config)
	return nil
}
