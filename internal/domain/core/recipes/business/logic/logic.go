package logic

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/code"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/set"
	"log"
	"math/rand"
	"time"
)

const (
	minIngredientNumber = 1
	maxIngredientNumber = 30
)

func NewRecipeSetGenerator(store port.RecipeStore) port.RecipeSetGenerator {
	return &recipeSetGenerator{RecipeStore: store}
}

type recipeSetGenerator struct {
	port.RecipeStore
}

func (r recipeSetGenerator) GenerateRecipeSet(ctx context.Context, recipesNumber uint32, ingredientsNumber uint32) error {
	recipes := make([]*model.Recipe, 0)

	if recipesNumber < 1 {
		return nil
	}

	if ingredientsNumber < minIngredientNumber {
		return code.New(consts.InvalidMin, fmt.Sprintf("recipes needs at least %d ingredient", minIngredientNumber))
	}

	if ingredientsNumber > maxIngredientNumber {
		return code.New(consts.InvalidMax, fmt.Sprintf("recipes needs at least %d ingredient", maxIngredientNumber))
	}

	log.Printf("%d recipes will generated\n", ingredientsNumber)
	for recipesNumber > 0 {
		now := time.Now().UnixNano()

		ingredientsSet := r.generateIngredientSet(ingredientsNumber)

		recipe := model.Recipe{
			ID:          now,
			Name:        fmt.Sprintf("Recipe #%d", now),
			Ingredients: ingredientsSet.Slice(),
		}

		recipes = append(recipes, &recipe)
		recipesNumber--
	}

	return r.CreateRecipes(ctx, recipes...)
}

func (recipeSetGenerator) generateIngredientSet(length uint32) set.Set[int64] {
	ingredients := make(set.Set[int64])

	for length > 0 {
		ingredient := rand.Int63()

		if ingredients.Has(ingredient) {
			continue
		}

		ingredients.Put(ingredient)
		length--
	}

	return ingredients
}
