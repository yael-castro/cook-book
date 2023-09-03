package business

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/pkg/errors/code"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
	"github.com/yael-castro/cb-search-engine-api/pkg/set"
	"strconv"
	"strings"
)

func NewRecipeAdder(creator RecipesCreator) RecipesAdder {
	if creator == nil {
		panic("nil dependency")
	}

	return &recipeAdder{
		creator: creator,
	}
}

type recipeAdder struct {
	creator RecipesCreator
}

func (r recipeAdder) AddRecipes(ctx context.Context, recipes ...*Recipe) error {
	for _, recipe := range recipes {
		if err := recipe.Validate(); err != nil {
			return err
		}
	}

	return r.creator.CreateRecipes(ctx, recipes...)
}

func NewRecipeGenerator(maker RecipesMaker, creator RecipesCreator) RecipesGenerator {
	switch any(nil) {
	case maker, creator:
		panic("missing settings")
	}

	return &recipesGenerator{
		maker:   maker,
		creator: creator,
	}
}

type recipesGenerator struct {
	maker   RecipesMaker
	creator RecipesCreator
}

func (r recipesGenerator) GenerateRecipes(ctx context.Context, recipesNumber uint32, ingredientsNumber uint32) error {
	if ingredientsNumber < minIngredientNumber {
		return code.New(InvalidMin, fmt.Sprintf("recipes needs at least %d ingredients", minIngredientNumber))
	}

	if ingredientsNumber > maxIngredientNumber {
		return code.New(InvalidMax, fmt.Sprintf("recipes can only have a maximum of %d ingredients", maxIngredientNumber))
	}

	recipes, err := r.maker.MakeRecipes(recipesNumber, ingredientsNumber)
	if err != nil {
		return err
	}

	return r.creator.CreateRecipes(ctx, recipes...)
}

// NewRecipesFinder builds a materialization for the port.RecipesSearcher interface
func NewRecipesFinder(searcher RecipesSearcher) RecipesFinder {
	if searcher == nil {
		panic("nil dependency")
	}

	return &recipesFinder{
		searcher: searcher,
	}
}

type recipesFinder struct {
	searcher RecipesSearcher
}

func (r recipesFinder) FindRecipes(ctx context.Context, str string, pagination *pagination.Pagination) ([]*Recipe, error) {
	if str == "" {
		return nil, code.New(MissingIngredientIdentifiers, "missing ingredients to make a recipe filter")
	}

	filter := &RecipeFilter{
		Pagination:  pagination,
		Ingredients: make(set.Set[int64]),
	}

	ingredients := strings.SplitN(str, ",", 10)

	for _, v := range ingredients {
		ingredient, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, code.New(InvalidIngredientID, fmt.Sprintf("ingredient id '%s' is not valid number", v))
		}

		filter.Ingredients.Put(ingredient)
	}

	return r.searcher.SearchRecipes(ctx, filter)
}
