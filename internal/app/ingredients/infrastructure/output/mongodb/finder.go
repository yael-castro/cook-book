package mongodb

import (
	"context"
	"github.com/yael-castro/cook-book/internal/app/ingredients/business"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewIngredientsFinder(collection *mongo.Collection, logger *log.Logger) business.IngredientsFinder {
	return ingredientsFinder{
		logger:     logger,
		collection: collection,
	}
}

type ingredientsFinder struct {
	logger     *log.Logger
	collection *mongo.Collection
}

func (i ingredientsFinder) FindIngredients(ctx context.Context, ingredientFilter *business.IngredientFilter) (business.Ingredients, error) {
	filter := NewIngredientFilter(*ingredientFilter)

	i.logger.Printf("FindIngredients: %+v\n", filter)

	cursor, err := i.findIngredients(ctx, filter)
	if err != nil {
		return nil, err
	}

	ingredients := make([]business.Ingredient, 0, filter.Size)

	for cursor.Next(ctx) {
		ingredient := Ingredient{}

		err = cursor.Decode(&ingredient)
		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient.ToBusinessModel())
	}

	return ingredients, nil
}

func (i ingredientsFinder) findIngredients(ctx context.Context, filter IngredientFilter) (*mongo.Cursor, error) {
	if filter.Random {
		opts := options.Aggregate().SetBatchSize(int32(filter.Size))

		return i.collection.Aggregate(ctx, filter.Pipeline(), opts)
	}

	// Establishing the limit of results and skip documents
	sorted := options.Find().SetSort(bson.D{
		{Key: "_id", Value: 1},
	})

	opts := []*options.FindOptions{
		sorted.SetSkip(int64(filter.Page * filter.Size)),
		sorted.SetLimit(int64(filter.Size)),
	}

	return i.collection.Find(ctx, filter.Document(), opts...) // TODO: find the better way to build a filter
}
