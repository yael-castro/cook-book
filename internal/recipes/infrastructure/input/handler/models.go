package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/input/handler"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
)

func NewRecipe(recipe *business.Recipe) *Recipe {
	return &Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Description: recipe.Description,
		Ingredients: handler.NewIngredients(recipe.Ingredients),
	}
}

type Recipe struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
}

func (r *Recipe) ToBModel() *business.Recipe {
	return &business.Recipe{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Ingredients: handler.BusinessIngredients(r.Ingredients),
	}
}

// RecipePage is a data transfer object for the recipe page
type RecipePage struct {
	Recipes []*Recipe `json:"recipes"`
	Total   uint64    `json:"total"`
}

type Ingredient = handler.Ingredient