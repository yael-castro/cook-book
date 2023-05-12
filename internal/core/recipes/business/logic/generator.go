package logic

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/errors/code"
)

const (
	minIngredientNumber = 1
	maxIngredientNumber = 30
)

func NewRecipeSetGenerator(generator port.RecipesGenerator, store port.RecipeStore) port.RecipeSetGenerator {
	return &recipeSetGenerator{
		RecipesGenerator: generator,
		RecipeStore:      store,
	}
}

type recipeSetGenerator struct {
	port.RecipesGenerator
	port.RecipeStore
}

func (r recipeSetGenerator) GenerateRecipeSet(ctx context.Context, recipesNumber uint32, ingredientsNumber uint32) error {
	if ingredientsNumber < minIngredientNumber {
		return code.New(consts.InvalidMin, fmt.Sprintf("recipes needs at least %d ingredients", minIngredientNumber))
	}

	if ingredientsNumber > maxIngredientNumber {
		return code.New(consts.InvalidMax, fmt.Sprintf("recipes can only have a maximum of %d ingredients", maxIngredientNumber))
	}

	recipes, err := r.GenerateRecipes(recipesNumber, ingredientsNumber)
	if err != nil {
		return err
	}

	return r.CreateRecipes(ctx, recipes...)
}
