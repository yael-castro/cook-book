package business

import (
	"context"
)

// Ports for drive adapters
type (
	RecipesSearcher interface {
		SearchRecipes(context.Context, *RecipeFilter) ([]*Recipe, error)
	}

	RecipesGenerator interface {
		GenerateRecipes(context.Context, GenerateRecipes) error
	}

	RecipesCreator interface {
		CreateRecipes(context.Context, ...*Recipe) error
	}
)

// Ports for driven adapters
type (
	RecipesFinder interface {
		FindRecipes(context.Context, *RecipeFilter) ([]*Recipe, error)
	}

	RecipesWriter interface {
		WriteRecipes(context.Context, GenerateRecipes) ([]*Recipe, error) // TODO: evaluate struct naming
	}

	RecipesSaver interface {
		SaveRecipes(context.Context, ...*Recipe) error
	}
)

var _ RecipesSaver = RecipesSaverFunc(nil)

type RecipesSaverFunc func(context.Context, ...*Recipe) error

func (f RecipesSaverFunc) SaveRecipes(ctx context.Context, recipe ...*Recipe) error {
	return f(ctx, recipe...)
}
