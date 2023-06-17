package logic

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/pkg/errors/code"
)

const (
	minIngredientNumber = 1
	maxIngredientNumber = 30
)

func NewRecipeSetGenerator(maker port.RecipesMaker, creator port.RecipesCreator) port.RecipesGenerator {
	switch any(nil) {
	case maker, creator:
		panic("nil dependency")
	}

	return &recipesGenerator{
		RecipesMaker:   maker,
		RecipesCreator: creator,
	}
}

type recipesGenerator struct {
	port.RecipesMaker
	port.RecipesCreator
}

func (r recipesGenerator) GenerateRecipes(ctx context.Context, recipesNumber uint32, ingredientsNumber uint32) error {
	if ingredientsNumber < minIngredientNumber {
		return code.New(consts.InvalidMin, fmt.Sprintf("recipes needs at least %d ingredients", minIngredientNumber))
	}

	if ingredientsNumber > maxIngredientNumber {
		return code.New(consts.InvalidMax, fmt.Sprintf("recipes can only have a maximum of %d ingredients", maxIngredientNumber))
	}

	recipes, err := r.MakeRecipes(recipesNumber, ingredientsNumber)
	if err != nil {
		return err
	}

	return r.CreateRecipes(ctx, recipes...)
}
