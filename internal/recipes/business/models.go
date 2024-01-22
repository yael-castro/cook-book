package business

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business"
	"github.com/yael-castro/cb-search-engine-api/pkg/set"
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
		return fmt.Errorf("%w: %d is not a valid page size", ErrInvalidPageSize, r.Size)
	case r.Ingredients.Len() == 0:
		return fmt.Errorf("%w: missing ingredients to perform a recipe search", ErrInvalidIngredients)
	}

	for ingredientID := range r.Ingredients {
		if ingredientID <= 0 {
			return fmt.Errorf("%w: %d is not a valid ingredient id", ErrInvalidIngredients, ingredientID)
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

const (
	minRecipes     = 1
	minIngredients = 1
	maxIngredients = 30
)

type GenerateRecipes struct {
	Recipes, Ingredients uint32
}

func (g GenerateRecipes) Validate() error {
	switch {
	case g.Recipes < minRecipes:
		return fmt.Errorf("%w: missing at least recipe", ErrInvalidRecipes)
	case g.Ingredients < minIngredients:
		return fmt.Errorf("%w: recipes needs at least %d ingredients", ErrInvalidIngredients, minIngredients)
	case g.Ingredients > maxIngredients:
		return fmt.Errorf("%w: recipes can only have a maximum of %d ingredients", ErrInvalidIngredients, maxIngredients)
	}

	return nil
}

// Aliases
type (
	Ingredient             = business.Ingredient
	NutritionalInformation = business.NutritionalInformation
)
