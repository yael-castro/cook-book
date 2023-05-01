package dto

import (
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/model"
)

func ToModelRecipe(recipe *Recipe) *model.Recipe {
	return (*model.Recipe)(recipe)
}

type Recipe struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Ingredients []int64 `json:"ingredients"`
}
