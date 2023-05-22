package dto

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/input/dto"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
)

func NewRecipe(recipe *model.Recipe) *Recipe {
	ingredients := make([]*Ingredient, 0, len(recipe.Ingredients))

	for _, ingredient := range recipe.Ingredients {
		ingredients = append(ingredients, dto.NewIngredient(ingredient))
	}

	return &Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Description: recipe.Description,
		Ingredients: ingredients,
	}
}

func ToModelRecipe(recipe *Recipe) *model.Recipe {
	ingredients := make([]*model.Ingredient, 0, len(recipe.Ingredients))

	for _, ingredient := range recipe.Ingredients {
		ingredients = append(ingredients, dto.ToModelIngredient(ingredient))
	}

	return &model.Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Description: recipe.Description,
		Ingredients: ingredients,
	}
}
