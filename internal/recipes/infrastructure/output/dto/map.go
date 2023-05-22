package dto

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/output/dto"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
)

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
