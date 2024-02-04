package business

import "strconv"

// Supported values for IngredientError
const (
	_ IngredientError = iota
	ErrInvalidID
	ErrInvalidName
	ErrInvalidDescription
	ErrInvalidCategory
	ErrInvalidFilter
)

// IngredientError defines an error related to a recipe error
type IngredientError uint8

// Error returns the string value of IngredientError
func (r IngredientError) Error() string {
	return "ingredient:" + strconv.FormatUint(uint64(r), 10)
}
