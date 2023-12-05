package output

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
)

func BusinessRecipe(recipe *Recipe) *business.Recipe {
	ingredients := make([]business.Ingredient, 0, len(recipe.Ingredients))

	for _, ingredient := range recipe.Ingredients {
		ingredients = append(ingredients, output.BusinessIngredient(ingredient))
	}

	return &business.Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Description: recipe.Description,
		Ingredients: ingredients,
	}
}

// NewRecipe builds an instance of *Recipe based on the *model.Recipe
func NewRecipe(recipe *business.Recipe) *Recipe {
	ingredients := make([]Ingredient, 0, len(recipe.Ingredients))

	for _, ingredient := range recipe.Ingredients {
		ingredients = append(ingredients, output.NewIngredient(ingredient))
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
	ID          int64        `bson:"_id"`
	Name        string       `bson:",omitempty"`
	Description string       `bson:",omitempty"`
	Ingredients []Ingredient `bson:",omitempty"`
}

type Ingredient = output.Ingredient
