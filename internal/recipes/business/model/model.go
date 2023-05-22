package model

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business/model"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
	"github.com/yael-castro/cb-search-engine-api/pkg/set"
)

// RecipeFilter filter for paginated recipe searches
type RecipeFilter struct {
	*pagination.Pagination
	// Ingredients are the ingredients used to perform a recipe search
	Ingredients set.Set[int64]
}

// Recipe kitchen recipe data
type Recipe struct {
	ID                int64
	Name, Description string
	Ingredients       []*Ingredient
}

func (r Recipe) IsValid() (bool, error) {
	return false, nil
}

// Ingredient alias for model.Ingredient
type Ingredient = model.Ingredient

// NutritionalInformation alias for model.NutritionalInformation
type NutritionalInformation = model.NutritionalInformation
