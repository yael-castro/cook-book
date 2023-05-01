package logic

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/core/searches/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/core/searches/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/core/searches/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/code"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/pagination"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/set"
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

func (r recipeSearcher) SearchRecipe(ctx context.Context, ingredients string, pagination *pagination.Pagination) ([]*model.Recipe, error) {
	search := &model.RecipeFilter{
		Pagination:  pagination,
		Ingredients: make(set.Set[int64]),
	}

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
