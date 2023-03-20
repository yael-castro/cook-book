package port

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/model"
)

type (
	// RecipeSetGenerator defines a application port to generate recipes and then save it
	RecipeSetGenerator interface {
		// GenerateRecipeSet generates a recipe set
		GenerateRecipeSet(context.Context, uint32, uint32) error
	}

	// RecipeStore defines a infrastructure port to manage the recipe storage
	RecipeStore interface {
		// CreateRecipes creates many recipe records until the context is not canceled
		CreateRecipes(context.Context, ...*model.Recipe) error
		// UpdateRecipe updates a recipe from the storage by id
		UpdateRecipe(context.Context, uint64) error
		// DeleteRecipe removes a recipe from the storage by id
		DeleteRecipe(context.Context, uint64) error
	}
)
