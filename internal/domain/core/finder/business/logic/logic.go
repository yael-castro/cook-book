package logic

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/port"
	recipes "github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/code"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/pagination"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/set"
	"log"
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
	search := &model.RecipeFilter{Pagination: pagination, Ingredients: make(set.Set[int64])}

	if ingredients == "" {
		return nil, code.New(consts.MissingIngredientIdentifiers, "missing ingredients to make a recipe search")
	}

	ingredientSlice := strings.SplitN(ingredients, ",", 10)

	for _, v := range ingredientSlice {
		ingredient, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Println(err)
			return nil, code.New(consts.InvalidIngredientID, fmt.Sprintf("ingredient id '%s' is not valid number", v))
		}

		search.Ingredients.Put(ingredient)
	}

	return r.RecipeFinder.FindRecipe(ctx, search)
}
