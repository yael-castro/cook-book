package output

import (
	"context"
	"errors"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/pkg/connection"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestRecipesCreator_CreateRecipes(t *testing.T) {
	cases := [...]struct {
		recipes     []*business.Recipe
		expectedErr error
	}{
		{
			recipes: []*business.Recipe{
				{
					ID:          1_022,
					Name:        "Recipe",
					Description: "<Insert ingredient description>",
					Ingredients: []*business.Ingredient{
						{
							ID:          1,
							Name:        "Ingredient",
							Description: "<Insert ingredient description>",
						},
					},
				},
			},
		},
	}

	database, err := connection.NewMongoDatabase(os.Getenv("MONGO_DSN"), os.Getenv("MONGO_DB"))
	if err != nil {
		t.Fatal(err)
	}

	creator := NewRecipeCreator(database, log.Default())

	for i, v := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := creator.CreateRecipes(context.TODO(), v.recipes...)
			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' got '%v'", v.expectedErr, err)
			}

			t.Log("Success!")
		})
	}
}
