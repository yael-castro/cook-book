package container

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	ingredientsbusiness "github.com/yael-castro/cb-search-engine-api/internal/ingredients/business"
	ingredientshandler "github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/input/handler"
	ingredientsout "github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/output"
	recipesbusiness "github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	recipeshandler "github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input/handler"
	recipesout "github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var GitCommit = ""

func Inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case *echo.Echo:
		return injectHandler(ctx, a)
	case *mongo.Database:
		return injectMongoDatabase(ctx, a)
	}

	return fmt.Errorf("type \"%T\" is not supported", a)
}

func injectHandler(ctx context.Context, e *echo.Echo) (err error) {
	// External dependencies
	var db mongo.Database

	if err = Inject(ctx, &db); err != nil {
		return
	}

	logger := log.Default()

	// MongoDB collections
	recipesCollection := db.Collection("recipes")

	// Driven adapters
	recipeSaver := recipesout.NewRecipesSaver(&db, logger)
	recipeFinder := recipesout.NewRecipesFinder(recipesCollection)
	ingredientFinder := ingredientsout.NewIngredientsFinder()

	// Ports for primary adapters
	recipeSearcher := recipesbusiness.NewRecipesSearcher(recipeFinder)
	recipeCreator := recipesbusiness.NewRecipeCreator(recipeSaver)
	ingredientSearcher := ingredientsbusiness.NewIngredientsSearcher(ingredientFinder)

	// Setting drive adapters
	e.POST(
		"/v1/recipes",
		recipeshandler.PostRecipes(recipeCreator),
	)

	e.GET(
		"/v1/recipes",
		recipeshandler.GetRecipes(recipeSearcher),
	)

	e.GET(
		"/v1/ingredients",
		ingredientshandler.GetIngredients(ingredientSearcher),
	)

	// Setting http error handler
	e.HTTPErrorHandler = ingredientshandler.ErrorHandler(
		recipeshandler.ErrorHandler(
			e.HTTPErrorHandler,
		),
	)

	return
}

func injectMongoDatabase(ctx context.Context, database *mongo.Database) (err error) {
	dsn, dbName := os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB")

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	*database = *mongoClient.Database(dbName)

	err = database.Client().Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	return
}
