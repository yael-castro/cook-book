package dto

import "github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/model"

// NewRecipe builds an instance of *Recipe based on the *model.Recipe
func NewRecipe(recipe *model.Recipe) *Recipe {
	return (*Recipe)(recipe)
}

// Recipe kitchen recipe data
type Recipe struct {
	ID          int64   `bson:"_id"`
	Name        string  `bson:",omitempty"`
	Description string  `bson:",omitempty"`
	Ingredients []int64 `bson:",omitempty"`
}
