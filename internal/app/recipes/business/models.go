package business

import (
	"fmt"
	"github.com/yael-castro/cook-book/internal/app/ingredients/business"
	"github.com/yael-castro/cook-book/pkg/set"
)

// RecipeFilter filter for paginated recipe searches
type RecipeFilter struct {
	Page, Size, Total uint64
	// Ingredients are the ingredient IDs used to perform a recipe search
	Ingredients set.Set[int64]
}

func (r RecipeFilter) Validate() error {
	const minSize, maxSize = 1, 25

	switch {
	case r.Size < minSize || r.Size > maxSize:
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
	Ingredients       Ingredients
}

func (r *Recipe) Validate() (err error) {
	if r == nil {
		return fmt.Errorf("%w: missing recipe data", ErrInvalidRecipe)
	}

	if len(r.Name) == 0 {
		return fmt.Errorf("%w: missing recipe name", ErrInvalidRecipe)
	}

	return r.Ingredients.Validate()
}

// GenerateRecipes contains the parameters used generate recipes
type GenerateRecipes struct {
	Recipes, Ingredients uint32
}

func (g GenerateRecipes) Validate() error {
	const (
		minRecipes     = 1
		minIngredients = 1
		maxIngredients = 30
	)

	switch {
	case g.Recipes < minRecipes:
		return fmt.Errorf("%w: missing at least %d recipe(s)", ErrInvalidRecipes, minRecipes)
	case g.Ingredients < minIngredients:
		return fmt.Errorf("%w: recipes needs at least %d ingredient(s)", ErrInvalidIngredients, minIngredients)
	case g.Ingredients > maxIngredients:
		return fmt.Errorf("%w: recipes can only have a maximum of %d ingredient(s)", ErrInvalidIngredients, maxIngredients)
	}

	return nil
}

// Aliases
type (
	Ingredient             = business.Ingredient
	Ingredients            = business.Ingredients
	NutritionalInformation = business.NutritionalInformation
)
