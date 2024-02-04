package output

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business"
)

func NewIngredientsFinder() business.IngredientsFinder {
	return ingredientsFinder{}
}

type ingredientsFinder struct{}

func (i ingredientsFinder) FindIngredients(ctx context.Context, filter *business.IngredientFilter) ([]business.Ingredient, error) {
	// TODO: implement the logic to search ingredients based on the filter
	return []business.Ingredient{}, nil
}
