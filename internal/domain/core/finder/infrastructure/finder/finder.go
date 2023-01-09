package finder

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/port"
	recipes "github.com/yael-castro/cb-search-engine-api/internal/domain/core/recipes/business/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s recipeFinder) FindRecipe(ctx context.Context, search *model.RecipeFilter) (slice []*recipes.Recipe, err error) {
	sorted := options.Find().SetSort(bson.D{{"_id", 1}}) // Use pool pattern?

	opts := []*options.FindOptions{sorted.SetSkip(int64(search.Start())), sorted.SetLimit(int64(search.Limit()))}

	searchFilter := bson.D{{"ingredients", bson.D{{"$all", search.Ingredients.Slice()}}}}

	totalResults, err := s.RecipeCollection.CountDocuments(ctx, searchFilter)
	if err != nil {
		return
	}

	if totalResults == 0 {
		return
	}

	search.SetTotalResults(uint64(totalResults))

	cursor, err := s.RecipeCollection.Find(ctx, searchFilter, opts...)
	if err != nil {
		return nil, err
	}

	// Avoid the method cursor.All() because it uses reflection to iterate the cursor values
	for cursor.Next(ctx) {
		recipe := recipes.Recipe{}

		err = cursor.Decode(&recipe)
		if err != nil {
			return
		}

		slice = append(slice, &recipe)
	}

	return
}
