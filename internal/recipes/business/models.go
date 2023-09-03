package business

import (
	"errors"
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business"
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

func (r *Recipe) Validate() (err error) {
	if r == nil {
		return errors.New("missing recipe data")
	}

	switch {
	case r.Name == "":
		return errors.New("missing recipe name")
	case r.Description == "":
		return errors.New("missing recipe description")
	}

	for _, ingredient := range r.Ingredients {
		if err = ingredient.Validate(); err != nil {
			return
		}
	}

	return
}

// Ingredient alias for model.Ingredient
type Ingredient = business.Ingredient

// NutritionalInformation alias for model.NutritionalInformation
type NutritionalInformation = business.NutritionalInformation
