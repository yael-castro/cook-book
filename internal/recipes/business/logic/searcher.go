package logic

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/pkg/errors/code"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
	"github.com/yael-castro/cb-search-engine-api/pkg/set"

	"strconv"
	"strings"
)

// NewRecipesFinder builds a materialization for the port.RecipesSearcher interface
func NewRecipesFinder(searcher port.RecipesSearcher) port.RecipesFinder {
	if searcher == nil {
		panic("nil dependency")
	}

	return &recipesSearcher{
		RecipesSearcher: searcher,
	}
}

type recipesSearcher struct {
	port.RecipesSearcher
}

func (r recipesSearcher) FindRecipes(ctx context.Context, str string, pagination *pagination.Pagination) ([]*model.Recipe, error) {
	filter := &model.RecipeFilter{
		Pagination:  pagination,
		Ingredients: make(set.Set[int64]),
	}

	if str == "" {
		return nil, code.New(consts.MissingIngredientIdentifiers, "missing ingredients to make a recipe filter")
	}

	ingredients := strings.SplitN(str, ",", 10)

	for _, v := range ingredients {
		ingredient, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, code.New(consts.InvalidIngredientID, fmt.Sprintf("ingredient id '%s' is not valid number", v))
		}

		filter.Ingredients.Put(ingredient)
	}

	return r.SearchRecipes(ctx, filter)
}
