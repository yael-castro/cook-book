package business

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business"
	"github.com/yael-castro/cb-search-engine-api/pkg/set"
	"strconv"
)

// RecipeFilter filter for paginated recipe searches
type RecipeFilter struct {
	Page, Size, Total uint64
	// Ingredients are the ingredient IDs used to perform a recipe search
	Ingredients set.Set[int64]
}

func (r RecipeFilter) Validate() error {
	switch {
	case r.Size < 1 || r.Size > 25:
		return ErrInvalidPageSize
	case r.Ingredients.Len() == 0:
		return fmt.Errorf("%w: missing ingredients to perform a recipe search", ErrInvalidIngredients)
	}

	for ingredientID := range r.Ingredients {
		if ingredientID <= 0 {
			return fmt.Errorf("%w: invalid ingredient id", ErrInvalidIngredients)
		}
	}

	return nil
}

type Recipe struct {
	ID                int64
	Name, Description string
	Ingredients       []Ingredient
}

func (r *Recipe) Validate() (err error) {
	if r == nil {
		return fmt.Errorf("%w: missing recipe data", ErrInvalidRecipe)
	}

	switch {
	case len(r.Name) == 0:
		return fmt.Errorf("%w: missing recipe name", ErrInvalidRecipe)
	}

	for _, ingredient := range r.Ingredients {
		if err = ingredient.Validate(); err != nil {
			return
		}
	}

	return
}

// Supported values for RecipeError
const (
	_ RecipeError = iota
	ErrInvalidRecipe
	ErrInvalidPageSize
	ErrInvalidIngredients
	ErrInvalidIngredientID
)

// RecipeError defines an error related to a recipe error
type RecipeError uint8

// Error returns the string value of RecipeError
func (r RecipeError) Error() string {
	return "recipes:" + strconv.FormatUint(uint64(r), 10)
}

// Aliases
type (
	Ingredient             = business.Ingredient
	NutritionalInformation = business.NutritionalInformation
)
