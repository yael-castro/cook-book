//go:build http

package mongodb_test

import (
	"context"
	"errors"
	"github.com/yael-castro/cook-book/internal/app/recipes/business"
	"github.com/yael-castro/cook-book/internal/app/recipes/infrastructure/output"
	"github.com/yael-castro/cook-book/internal/container"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
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
					Ingredients: []business.Ingredient{
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

	ctx := context.TODO()
	log.SetOutput(io.Discard)

	var mongoDB mongo.Database

	if err := container.Inject(ctx, &mongoDB); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		mongoDB.Client().Disconnect(ctx)
	})

	saver := output.NewRecipesSaver(&mongoDB, log.Default())

	for i, v := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := saver.SaveRecipes(ctx, v.recipes...)
			if !errors.Is(err, v.expectedErr) {
				t.Fatalf("expected error '%v' got '%v'", v.expectedErr, err)
			}

			t.Cleanup(func() {
				mongoDB.Collection("recipes").DeleteMany(ctx, nil)
			})

			t.Log("Success!")
		})
	}
}
