package port

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
)

// Ports for drive adapters
type (
	RecipesFinder interface {
		FindRecipes(context.Context, string, *pagination.Pagination) ([]*model.Recipe, error)
	}

	RecipesGenerator interface {
		GenerateRecipes(context.Context, uint32, uint32) error
	}

	RecipesAdder interface {
		AddRecipes(context.Context, ...*model.Recipe) error
	}
)

// Ports for driven adapters
type (
	RecipesSearcher interface {
		SearchRecipes(context.Context, *model.RecipeFilter) ([]*model.Recipe, error)
	}

	RecipesMaker interface {
		MakeRecipes(uint32, uint32) ([]*model.Recipe, error)
	}

	RecipesCreator interface {
		CreateRecipes(context.Context, ...*model.Recipe) error
	}
)
