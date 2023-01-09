package port

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/pagination"
)
import recipes "github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/model"

type (
	// RecipeSearcher defines a primary port for recipe lookups
	RecipeSearcher interface {
		// SearchRecipe receives a string to parse and based on it perform a paginated search for recipes
		SearchRecipe(context.Context, string, *pagination.Pagination) ([]*recipes.Recipe, error)
	}

	// RecipeFinder defines a secondary port to find recipes
	RecipeFinder interface {
		// FindRecipe search recipes based on the instance of *model.RecipeFilter
		FindRecipe(context.Context, *model.RecipeFilter) ([]*recipes.Recipe, error)
	}
)
