package output_test

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/cmd/server/container"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/pkg/set"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkRecipeFinder_FindRecipe(b *testing.B) {
	ingredientsFunc := func(n int) set.Set[int64] {
		ingredients := make(set.Set[int64])

		for len(ingredients) < n {
			ingredients.Add(rand.Int63())
		}

		return ingredients
	}

	cases := [...]struct {
		ctx         context.Context
		ingredients set.Set[int64]
	}{
		{
			ctx:         context.TODO(),
			ingredients: ingredientsFunc(5),
		},
		{
			ctx:         context.TODO(),
			ingredients: ingredientsFunc(10),
		},
		{
			ctx:         context.TODO(),
			ingredients: ingredientsFunc(15),
		},
		{
			ctx:         context.TODO(),
			ingredients: ingredientsFunc(20),
		},
	}

	log.SetOutput(io.Discard)
	var mongoDB mongo.Database

	if err := container.Inject(context.TODO(), &mongoDB); err != nil {
		b.Fatal(err)
	}

	recipeCollection := mongoDB.Collection("recipes")
	finder := output.NewRecipesFinder(recipeCollection)

	for i, c := range cases {
		b.Log("INGREDIENT NUMBER:", len(c.ingredients))
		b.Run(strconv.Itoa(i), func(b *testing.B) {

			for i := 0; i < b.N; i++ {
				_, err := finder.FindRecipes(c.ctx, &business.RecipeFilter{
					Page:        0,
					Size:        20,
					Ingredients: c.ingredients,
				})
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
