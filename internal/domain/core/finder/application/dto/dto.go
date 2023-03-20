package dto

import "github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/model"

// RecipePage is a data transfer object for the recipe page
type RecipePage struct {
	Items      any    `json:"items"`
	TotalPages uint64 `json:"totalPages"`
	TotalItems uint64 `json:"totalItems"`
}

func NewRecipe(recipe *model.Recipe) *Recipe {
	return (*Recipe)(recipe)
}

// Recipe is a data transfer object for the recipe data
type Recipe struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Ingredients []int64 `json:"ingredients"`
}
