package business

import "context"

func NewIngredientsSearcher(finder IngredientsFinder) IngredientSearcher {
	switch any(nil) {
	case finder:
		panic("missing parameters to initialize")
	}

	return ingredientsSearcher{
		finder: finder,
	}
}

type ingredientsSearcher struct {
	finder IngredientsFinder
}

func (i ingredientsSearcher) SearchIngredients(ctx context.Context, filter *IngredientFilter) (Ingredients, error) {
	if err := filter.Validate(); err != nil {
		return nil, err
	}

	return i.finder.FindIngredients(ctx, filter)
}
