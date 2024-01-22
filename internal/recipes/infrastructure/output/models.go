package output

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
)

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

func (r Recipe) ToBModel() *business.Recipe {
	ingredients := make([]business.Ingredient, 0, len(r.Ingredients))

	for _, ingredient := range r.Ingredients {
		ingredients = append(ingredients, output.BusinessIngredient(ingredient))
	}

	return &business.Recipe{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Ingredients: ingredients,
	}
}

type Ingredient = output.Ingredient
