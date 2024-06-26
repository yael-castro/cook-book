package mongodb

import (
	"github.com/yael-castro/cook-book/internal/app/ingredients/infrastructure/output/mongodb"
	"github.com/yael-castro/cook-book/internal/app/recipes/business"
)

// NewRecipe builds an instance of *Recipe based on the *model.Recipe
func NewRecipe(recipe *business.Recipe) *Recipe {
	ingredients := make(Ingredients, 0, len(recipe.Ingredients))

	for _, ingredient := range recipe.Ingredients {
		ingredients = append(ingredients, mongodb.NewIngredient(ingredient))
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
	ID          int64       `bson:"_id"`
	Name        string      `bson:",omitempty"`
	Description string      `bson:",omitempty"`
	Ingredients Ingredients `bson:",omitempty"`
}

func (r Recipe) ToBusinessModel() *business.Recipe {
	return &business.Recipe{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Ingredients: r.Ingredients.ToBusinessModel(),
	}
}

type Ingredients = mongodb.Ingredients
