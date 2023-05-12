package generator

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/set"
	"log"
	"math/rand"
	"time"
)

func NewRecipesGenerator() port.RecipesGenerator {
	return recipesGenerator{}
}

type recipesGenerator struct {
}

func (r recipesGenerator) GenerateRecipes(recipesNumber, ingredientsNumber uint32) ([]*model.Recipe, error) {
	recipes := make([]*model.Recipe, 0, recipesNumber)

	log.Printf("%d recipes will generated\n", ingredientsNumber)
	for recipesNumber > 0 {
		now := time.Now().UnixNano()

		ingredientsSet := r.generateIngredientSet(ingredientsNumber)

		recipe := model.Recipe{
			ID:          now,
			Name:        fmt.Sprintf("Recipe #%d", now),
			Description: "This is a test recipe",
			Ingredients: ingredientsSet.Slice(),
		}

		recipes = append(recipes, &recipe)
		recipesNumber--
	}

	return recipes, nil
}

func (r recipesGenerator) generateIngredientSet(length uint32) set.Set[int64] {
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