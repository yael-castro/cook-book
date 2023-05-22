package dto

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/input/dto"
)

type (
	// RecipePage is a data transfer object for the recipe page
	RecipePage struct {
		Items      any    `json:"items"`
		TotalPages uint64 `json:"totalPages"`
		TotalItems uint64 `json:"totalItems"`
	}

	Recipe struct {
		ID          int64         `json:"id"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Ingredients []*Ingredient `json:"ingredients"`
	}
)

type Ingredient = dto.Ingredient
