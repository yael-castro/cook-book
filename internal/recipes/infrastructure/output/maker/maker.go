package maker

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
	"log"
	"math/rand"
	"time"
)

func NewRecipesMaker() port.RecipesMaker {
	return &recipesMaker{}
}

type recipesMaker struct{}

func (r *recipesMaker) MakeRecipes(recipesNumber, ingredientsNumber uint32) ([]*model.Recipe, error) {
	recipes := make([]*model.Recipe, 0, recipesNumber)
	log.Printf("%d recipes will generated\n", ingredientsNumber)

	for recipesNumber > 0 {
		id := time.Now().UnixNano()

		recipe := model.Recipe{
			ID:          id,
			Name:        fmt.Sprintf("RECIPE #%d", id),
			Ingredients: r.generateIngredients(ingredientsNumber),
		}

		recipes = append(recipes, &recipe)
		recipesNumber--
	}

	return recipes, nil
}

func (*recipesMaker) generateIngredients(length uint32) []*model.Ingredient {
	ingredients := make([]*model.Ingredient, 0, length)

	for length > 0 {
		id := rand.Int63()

		ingredients = append(ingredients, &model.Ingredient{
			ID:   id,
			Name: fmt.Sprintf("INGREDIENT #%d", id),
			NutritionalInformation: &model.NutritionalInformation{
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
