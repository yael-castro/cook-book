//go:build integration

package mongodb_test

import (
	"context"
	"github.com/yael-castro/cook-book/internal/app/ingredients/business"
	"github.com/yael-castro/cook-book/internal/app/ingredients/infrastructure/output/mongodb"
	"github.com/yael-castro/cook-book/internal/container"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"testing"
)

func TestIngredientsFinder_FindIngredients(t *testing.T) {
	cases := [...]struct {
		ctx    context.Context
		filter *business.IngredientFilter
	}{
		{
			ctx: context.TODO(),
			filter: &business.IngredientFilter{
				Keyword: "tomato",
				Size:    1,
			},
		},
		{
			ctx: context.TODO(),
			filter: &business.IngredientFilter{
				Keyword: "INGREDIENT ",
				Size:    10,
			},
		},
		{
			ctx: context.TODO(),
			filter: &business.IngredientFilter{
				Random: true,
				Size:   5,
			},
		},
	}

	var db mongo.Database

	err := container.Inject(context.TODO(), &db)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = db.Client().Disconnect(context.Background())
	})

	ingredientsCollection := db.Collection("ingredients")

	finder := mongodb.NewIngredientsFinder(ingredientsCollection, log.Default())

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ingredients, err := finder.FindIngredients(c.ctx, c.filter)
			if err != nil { // TODO: evaluate handled errors
				t.Fatal(err)
			}

			for _, ingredient := range ingredients {
				t.Logf("Ingredient: %+v", ingredient)
			}
		})
	}
}
