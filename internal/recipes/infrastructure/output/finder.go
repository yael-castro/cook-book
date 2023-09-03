package output

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// NewRecipesSearcher builds an instance of the unique implementation for the port.RecipeFinder that use a MongoDB writes
func NewRecipesSearcher(collection *mongo.Collection) business.RecipesSearcher {
	if collection == nil {
		panic("missing MongoDB collection")
	}

	return &recipesSearcher{
		RecipeCollection: collection,
		logger:           log.Default(),
	}
}

type recipesSearcher struct {
	RecipeCollection *mongo.Collection
	logger           *log.Logger
}

func (s recipesSearcher) SearchRecipes(ctx context.Context, filter *business.RecipeFilter) (slice []*business.Recipe, err error) {
	s.logger.Printf("RECIPE FILTER: %+v\n", filter)

	// Establishing the limit of results and skip documents
	sorted := options.Find().SetSort(bson.D{
		{Key: "_id", Value: 1},
	})

	opts := []*options.FindOptions{
		sorted.SetSkip(int64(filter.Start())),
		sorted.SetLimit(int64(filter.Limit())),
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
	totalResults, err := s.RecipeCollection.CountDocuments(ctx, query)
	if err != nil {
		s.logger.Println(err)
		return
	}

	if totalResults < 1 {
		return
	}

	filter.SetTotalResults(uint64(totalResults))

	cursor, err := s.RecipeCollection.Find(ctx, query, opts...)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	// NOTE: avoid the method cursor.All() because it uses reflection to iterate the cursor values
	for cursor.Next(ctx) {
		recipe := &Recipe{}

		err = cursor.Decode(recipe)
		if err != nil {
			s.logger.Println(err)
			return
		}

		slice = append(slice, BusinessRecipe(recipe))
	}

	return
}
