package output

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"math/rand"
	"time"
)

func NewRecipesWriter() business.RecipesWriter {
	return recipesWriter{}
}

type recipesWriter struct{}

func (r recipesWriter) WriteRecipes(ctx context.Context, generate business.GenerateRecipes) ([]*business.Recipe, error) {
	recipes := make([]*business.Recipe, 0, generate.Recipes)

	for generate.Recipes > 0 {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		id := time.Now().UnixNano()

		recipe := business.Recipe{
			ID:          id,
			Name:        fmt.Sprintf("RECIPE #%d", id),
			Ingredients: r.generateIngredients(generate.Recipes),
		}

		recipes = append(recipes, &recipe)
		generate.Recipes--
	}

	return recipes, nil
}

func (recipesWriter) generateIngredients(length uint32) []business.Ingredient {
	ingredients := make([]business.Ingredient, 0, length)

	for length > 0 {
		id := rand.Int63()

		ingredients = append(ingredients, business.Ingredient{
			ID:   id,
			Name: fmt.Sprintf("INGREDIENT #%d", id),
			NutritionalInformation: business.NutritionalInformation{
				Fats:     1_000,
				Carbs:    1_000,
				Fiber:    1_000,
				Calories: 1_000,
				Proteins: 1_000,
			},
		})

		length--
	}

	return ingredients
}
