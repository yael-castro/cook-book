package storage

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/infrastructure/output/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// NewRecipeStore builds an adapter for the port.RecipeStore
func NewRecipeStore(collection *mongo.Collection) port.RecipeStore {
	return &recipe{Collection: collection}
}

type recipe struct {
	defaultOrdered bool
	Collection     *mongo.Collection
}

func (r recipe) CreateRecipes(ctx context.Context, recipes ...*model.Recipe) error {
	documents := make([]any, 0, len(recipes))

	for _, recipe := range recipes {
		log.Println("Recipe ID:", recipe.ID)
		documents = append(documents, dto.NewRecipe(recipe))
	}

	opts := &options.InsertManyOptions{Ordered: &r.defaultOrdered}

	// TODO: uses batches of 1000 to insert many records
	_, err := r.Collection.InsertMany(ctx, documents, opts)
	return err
}

func (recipe) UpdateRecipe(context.Context, *model.Recipe) error {
	//TODO implement me
	panic("implement me")
}

func (recipe) DeleteRecipe(context.Context, uint64) error {
	//TODO implement me
	panic("implement me")
}
