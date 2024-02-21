package handler

import (
	ingredientshandler "github.com/yael-castro/cook-book/internal/app/ingredients/infrastructure/input/handler"
	"github.com/yael-castro/cook-book/internal/app/recipes/business"
)

func NewRecipe(recipe *business.Recipe) *Recipe {
	return &Recipe{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Description: recipe.Description,
		Ingredients: ingredientshandler.NewIngredients(recipe.Ingredients),
	}
}

type Recipe struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
}

func (r *Recipe) ToBusiness() *business.Recipe {
	return &business.Recipe{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Ingredients: ingredientshandler.BusinessIngredients(r.Ingredients),
	}
}

// RecipePage is a data transfer object for the recipe page
type RecipePage struct {
	Recipes []*Recipe `json:"recipes"`
	Total   uint64    `json:"total"`
}

type Ingredient = ingredientshandler.Ingredient
