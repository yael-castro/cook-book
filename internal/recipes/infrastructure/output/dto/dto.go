package dto

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/output/dto"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
)

// NewRecipe builds an instance of *Recipe based on the *model.Recipe
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

// Recipe kitchen recipe data
type Recipe struct {
	ID          int64         `bson:"_id"`
	Name        string        `bson:",omitempty"`
	Description string        `bson:",omitempty"`
	Ingredients []*Ingredient `bson:",omitempty"`
}

type Ingredient = dto.Ingredient
