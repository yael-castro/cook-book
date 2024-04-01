package business

import "strconv"

// Supported values for RecipeError
const (
	_ RecipeError = iota
	// ErrInvalidRecipe indicates an error caused by a nil recipe
	ErrInvalidRecipe
	// ErrInvalidPageSize indicates an error caused by an invalid page size for request recipes
	ErrInvalidPageSize
	// ErrInvalidIngredientID indicates an error caused by an invalid ingredient id
	ErrInvalidIngredientID
	// ErrInvalidIngredients indicates an error caused by an incorrect number of ingredients
	ErrInvalidIngredients
	// ErrInvalidRecipes indicates an error caused by an incorrect number of recipes
	ErrInvalidRecipes
)

// RecipeError defines an error related to a recipe error
type RecipeError uint8

// Error returns the string value of RecipeError
func (r RecipeError) Error() string {
	return "recipes:" + strconv.FormatUint(uint64(r), 10)
}
