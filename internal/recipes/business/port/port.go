package port

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
)

// Ports for input adapters
type (
	// RecipeSearcher defines an application port for recipe lookups
	RecipeSearcher interface {
		// SearchRecipe receives a string to parse and based on it perform a paginated search for recipes
		SearchRecipe(context.Context, string, *pagination.Pagination) ([]*model.Recipe, error)
	}

	// RecipeSetGenerator defines a random recipe generator
	RecipeSetGenerator interface {
		// GenerateRecipeSet generates a random recipe set based on the parameters
		GenerateRecipeSet(context.Context, uint32, uint32) error
	}

	// RecipeManager defines the common operations to manage the recipe storage
	RecipeManager interface {
		CreateRecipes(context.Context, ...*model.Recipe) error
		UpdateRecipe(context.Context, *model.Recipe) error
		DeleteRecipe(context.Context, uint64) error
	}
)

// Ports for output adapters
type (
	// RecipeFinder defines a sdk port to find recipes
	RecipeFinder interface {
		// FindRecipe searches recipes based on the instance of *model.RecipeFilter
		FindRecipe(context.Context, *model.RecipeFilter) ([]*model.Recipe, error)
	}

	RecipesGenerator interface {
		GenerateRecipes(uint32, uint32) ([]*model.Recipe, error)
	}

	// RecipeStore defines the common operations to manage the recipe storage
	RecipeStore interface {
		CreateRecipes(context.Context, ...*model.Recipe) error
		UpdateRecipe(context.Context, *model.Recipe) error
		DeleteRecipe(context.Context, uint64) error
	}
)
