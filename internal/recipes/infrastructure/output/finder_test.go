package output

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/pkg/connection"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
	"github.com/yael-castro/cb-search-engine-api/pkg/set"
	"math/rand"
	"os"
	"testing"
)

func BenchmarkRecipeFinder_FindRecipe(b *testing.B) {
	mongoDB, err := connection.NewMongoDatabase(os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB"))
	if err != nil {
		b.Fatal(err)
	}

	recipeCollection := mongoDB.Collection("recipes")

	searcher := NewRecipesSearcher(recipeCollection)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		ingredients := make(set.Set[int64])

		for len(ingredients) < 20 {
			ingredients.Put(rand.Int63())
		}

		b.StartTimer()

		_, err = searcher.SearchRecipes(context.TODO(), &business.RecipeFilter{
			Pagination:  pagination.New("0", "20"),
			Ingredients: ingredients,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}
