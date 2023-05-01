package finder

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/core/searches/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/connection"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/pagination"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/set"
	"math/rand"
	"testing"
)

const (
	defaultDSN = "mongodb://localhost:27017"
	defaultDB  = "test"
)

func BenchmarkRecipeFinder_FindRecipe(b *testing.B) {
	mongoDB, err := connection.NewMongoDatabase(defaultDSN, defaultDB)
	if err != nil {
		b.Fatal(err)
	}

	recipeCollection := mongoDB.Collection("recipes")

	finder := NewRecipeFinder(recipeCollection)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		ingredients := make(set.Set[int64])

		for len(ingredients) < 20 {
			ingredients.Put(rand.Int63())
		}

		b.StartTimer()

		finder.FindRecipe(context.TODO(), &model.RecipeFilter{
			Pagination:  pagination.New("0", "20"),
			Ingredients: ingredients,
		})
	}
}
