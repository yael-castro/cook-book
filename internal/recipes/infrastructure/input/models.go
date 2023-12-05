package input

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/input"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
)

func BusinessRecipe(recipe *Recipe) *business.Recipe {
	return &business.Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Description: recipe.Description,
		Ingredients: input.BusinessIngredients(recipe.Ingredients),
	}
}

func NewRecipe(recipe *business.Recipe) *Recipe {
	return &Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Description: recipe.Description,
		Ingredients: input.NewIngredients(recipe.Ingredients),
	}
}

type Recipe struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
}

// RecipePage is a data transfer object for the recipe page
type RecipePage struct {
	Recipes []*Recipe `json:"recipes"`
	Total   uint64    `json:"total"`
}

type Ingredient = input.Ingredient
