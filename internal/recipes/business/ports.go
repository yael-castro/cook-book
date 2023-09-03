package business

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
)

// Ports for drive adapters
type (
	RecipesFinder interface {
		FindRecipes(context.Context, string, *pagination.Pagination) ([]*Recipe, error)
	}

	RecipesGenerator interface {
		GenerateRecipes(context.Context, uint32, uint32) error
	}

	RecipesAdder interface {
		AddRecipes(context.Context, ...*Recipe) error
	}
)

// Ports for driven adapters
type (
	RecipesSearcher interface {
		SearchRecipes(context.Context, *RecipeFilter) ([]*Recipe, error)
	}

	RecipesMaker interface {
		MakeRecipes(uint32, uint32) ([]*Recipe, error)
	}

	RecipesCreator interface {
		CreateRecipes(context.Context, ...*Recipe) error
	}
)

var _ RecipesCreator = RecipesCreatorFunc(nil)

type RecipesCreatorFunc func(context.Context, ...*Recipe) error

func (f RecipesCreatorFunc) CreateRecipes(ctx context.Context, recipe ...*Recipe) error {
	return f(ctx, recipe...)
}
