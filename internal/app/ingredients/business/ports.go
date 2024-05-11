package business

import "context"

// Ports for drive adapters
type (
	// IngredientSearcher defines a way to search ingredients by filter
	IngredientSearcher interface {
		// SearchIngredients searches ingredients by filter
		SearchIngredients(context.Context, *IngredientFilter) (Ingredients, error)
	}
)

// Ports for driven adapters
type (
	// IngredientsFinder defines a way to find an ingredients
	IngredientsFinder interface {
		// FindIngredients finds an ingredient by filter
		FindIngredients(context.Context, *IngredientFilter) (Ingredients, error)
	}
)
