package finder

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// NewRecipeFinder builds an instance of the unique implementation for the port.RecipeFinder that use a MongoDB storage
func NewRecipeFinder(collection *mongo.Collection) port.RecipeFinder {
	return &recipeFinder{
		RecipeCollection: collection,
	}
}

type recipeFinder struct {
	RecipeCollection *mongo.Collection
}

func (s recipeFinder) FindRecipe(ctx context.Context, filter *model.RecipeFilter) (slice []*model.Recipe, err error) {
	log.Printf("RECIPE FILTER: %+v\n", filter)

	sorted := options.Find().SetSort(bson.D{
		{Key: "_id", Value: 1},
	})

	log.Printf("START %d LIMIT %d\n", filter.Start(), filter.Limit())

	opts := []*options.FindOptions{
		sorted.SetSkip(int64(filter.Start())),
		sorted.SetLimit(int64(filter.Limit())),
	}

	// Transforms the
	ingredients := make([]bson.D, 0, len(filter.Ingredients))

	for ingredient := range filter.Ingredients {
		ingredients = append(ingredients, bson.D{
			{
				Key:   "$elemMatch",
				Value: bson.D{{Key: "_id", Value: ingredient}},
			},
		})
	}

	log.Printf("INGREDIENTS: %+v", ingredients)

	query := bson.D{
		{
			Key: "ingredients",
			Value: bson.D{
				{
					Key:   "$all",
					Value: ingredients,
				},
			},
		},
	}

	totalResults, err := s.RecipeCollection.CountDocuments(ctx, query)
	if err != nil {
		return
	}

	if totalResults < 1 {
		return
	}

	filter.SetTotalResults(uint64(totalResults))

	cursor, err := s.RecipeCollection.Find(ctx, query, opts...)
	if err != nil {
		return nil, err
	}

	// NOTE: avoid the method cursor.All() because it uses reflection to iterate the cursor values
	for cursor.Next(ctx) {
		recipe := &dto.Recipe{}

		err = cursor.Decode(recipe)
		if err != nil {
			return
		}

		slice = append(slice, dto.ToModelRecipe(recipe))
	}

	return
}
