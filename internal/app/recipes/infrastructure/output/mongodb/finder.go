package mongodb

import (
	"context"
	"github.com/yael-castro/cook-book/internal/app/recipes/business"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// NewRecipesFinder builds an instance of the unique implementation for the port.RecipeFinder that use a MongoDB writes
func NewRecipesFinder(collection *mongo.Collection) business.RecipesFinder {
	if collection == nil {
		panic("missing MongoDB collection")
	}

	return recipesFinder{
		recipeCollection: collection,
		logger:           log.Default(),
	}
}

type recipesFinder struct {
	recipeCollection *mongo.Collection
	logger           *log.Logger
}

func (s recipesFinder) FindRecipes(ctx context.Context, filter *business.RecipeFilter) (slice []*business.Recipe, err error) {
	s.logger.Printf("RECIPE FILTER: %+v\n", filter)

	// Establishing the limit of results and skip documents
	sorted := options.Find().SetSort(bson.D{
		{Key: "_id", Value: 1},
	})

	opts := []*options.FindOptions{
		sorted.SetSkip(int64(filter.Page * filter.Size)),
		sorted.SetLimit(int64(filter.Size)),
	}

	// Encoding ingredient IDs
	ingredients := make([]bson.D, 0, len(filter.Ingredients))

	for ingredient := range filter.Ingredients {
		ingredients = append(ingredients, bson.D{
			{
				Key:   "$elemMatch",
				Value: bson.D{{Key: "_id", Value: ingredient}},
			},
		})
	}

	// Building the MongoDB query
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

	// Counting documents
	totalResults, err := s.recipeCollection.CountDocuments(ctx, query)
	if err != nil {
		s.logger.Println(err)
		return
	}

	if totalResults < 1 {
		return
	}

	filter.Total = uint64(totalResults)

	cursor, err := s.recipeCollection.Find(ctx, query, opts...)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	// NOTE: avoid the method cursor.All() because it uses reflection to iterate the cursor values
	for cursor.Next(ctx) {
		recipe := Recipe{}

		err = cursor.Decode(&recipe)
		if err != nil {
			s.logger.Println(err)
			return
		}

		slice = append(slice, recipe.ToBusinessModel())
	}

	return
}
