package logic

import (
	"context"
	"errors"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/port"
)

func NewRecipeManager(manager port.RecipeManager) port.RecipeManager {
	return &recipeManager{
		RecipeStore: manager,
	}
}

type recipeManager struct {
	RecipeStore port.RecipeStore
}

func (r recipeManager) CreateRecipes(ctx context.Context, recipes ...*model.Recipe) error {
	for _, recipe := range recipes {
		if isValid, err := recipe.IsValid(); !isValid {
			return err
		}
	}

	return r.RecipeStore.CreateRecipes(ctx, recipes...)
}

func (r recipeManager) UpdateRecipe(ctx context.Context, recipe *model.Recipe) error {
	if isValid, err := recipe.IsValid(); !isValid {
		return err
	}

	return r.RecipeStore.UpdateRecipe(ctx, recipe)
}

func (r recipeManager) DeleteRecipe(ctx context.Context, recipeID uint64) error {
	if recipeID == 0 {
		return errors.New("missing recipe id")
	}

	return r.RecipeStore.DeleteRecipe(ctx, recipeID)
}
