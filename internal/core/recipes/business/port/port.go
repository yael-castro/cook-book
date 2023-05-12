package port

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/model"
)

// Ports for driving adapters
type (
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

// Ports for drive adapters
type (
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
