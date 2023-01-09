package logic

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/port"
	recipes "github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/pagination"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/set"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/wrong"
	"strconv"
	"strings"
)

// NewRecipeSearcher builds a materialization for the port.RecipeSearcher interface
func NewRecipeSearcher(finder port.RecipeFinder) port.RecipeSearcher {
	return recipeSearcher{RecipeFinder: finder}
}

type recipeSearcher struct {
	RecipeFinder port.RecipeFinder
}

func (r recipeSearcher) SearchRecipe(ctx context.Context, ingredients string, pagination *pagination.Pagination) ([]*recipes.Recipe, error) {
	search := &model.RecipeFilter{Pagination: pagination, Ingredients: make(set.Set[int32])}

	if ingredients == "" {
		return nil, wrong.Validation("missing ingredients to make a recipe search")
	}

	ingredientSlice := strings.SplitN(ingredients, ",", 10)

	for _, v := range ingredientSlice {
		ingredient, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, wrong.Validation(err.Error())
		}

		search.Ingredients.Put(int32(ingredient))
	}

	return r.RecipeFinder.FindRecipe(ctx, search)
}
