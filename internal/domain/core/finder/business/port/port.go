package port

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/pagination"
)

type (
	// RecipeSearcher defines an application port for recipe lookups
	RecipeSearcher interface {
		// SearchRecipe receives a string to parse and based on it perform a paginated search for recipes
		SearchRecipe(context.Context, string, *pagination.Pagination) ([]*model.Recipe, error)
	}

	// RecipeFinder defines a infrastructure port to find recipes
	RecipeFinder interface {
		// FindRecipe searches recipes based on the instance of *model.RecipeFilter
		FindRecipe(context.Context, *model.RecipeFilter) ([]*model.Recipe, error)
	}
)
