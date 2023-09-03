package output

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"math/rand"
	"time"
)

func NewRecipesMaker() business.RecipesMaker {
	return recipesMaker{}
}

type recipesMaker struct{}

func (r recipesMaker) MakeRecipes(recipesNumber, ingredientsNumber uint32) ([]*business.Recipe, error) {
	recipes := make([]*business.Recipe, 0, recipesNumber)

	for recipesNumber > 0 {
		id := time.Now().UnixNano()

		recipe := business.Recipe{
			ID:          id,
			Name:        fmt.Sprintf("RECIPE #%d", id),
			Ingredients: r.generateIngredients(ingredientsNumber),
		}

		recipes = append(recipes, &recipe)
		recipesNumber--
	}

	return recipes, nil
}

func (recipesMaker) generateIngredients(length uint32) []*business.Ingredient {
	ingredients := make([]*business.Ingredient, 0, length)

	for length > 0 {
		id := rand.Int63()

		ingredients = append(ingredients, &business.Ingredient{
			ID:   id,
			Name: fmt.Sprintf("INGREDIENT #%d", id),
			NutritionalInformation: &business.NutritionalInformation{
				Calories: 1_000,
				Fats:     1_000,
				Carbs:    1_000,
				Fiber:    1_000,
				Proteins: 1_000,
			},
		})

		length--
	}

	return ingredients
}
