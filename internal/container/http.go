//go:build http

package container

import (
	"context"
	"github.com/labstack/echo/v4"
	ingredientsbusiness "github.com/yael-castro/cook-book/internal/app/ingredients/business"
	ingredientshandler "github.com/yael-castro/cook-book/internal/app/ingredients/infrastructure/input/handler"
	ingredientsmongo "github.com/yael-castro/cook-book/internal/app/ingredients/infrastructure/output/mongodb"
	recipesbusiness "github.com/yael-castro/cook-book/internal/app/recipes/business"
	recipeshandler "github.com/yael-castro/cook-book/internal/app/recipes/infrastructure/input/handler"
	recipesmongo "github.com/yael-castro/cook-book/internal/app/recipes/infrastructure/output/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case *echo.Echo:
		return injectHandler(ctx, a)
	}

	return inject(ctx, a)
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
	ingredientsCollection := db.Collection("ingredients")

	// Driven adapters
	recipeSaver := recipesmongo.NewRecipesSaver(&db, logger)
	recipeFinder := recipesmongo.NewRecipesFinder(recipesCollection)
	ingredientFinder := ingredientsmongo.NewIngredientsFinder(ingredientsCollection, logger)

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
