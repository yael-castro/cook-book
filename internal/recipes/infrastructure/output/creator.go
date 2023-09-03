package output

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// NewRecipeCreator builds an adapter for the port.RecipesCreator
func NewRecipeCreator(recipesCollection *mongo.Collection) business.RecipesCreator {
	if recipesCollection == nil {
		panic("missing MongoDB collection")
	}

	return &recipesCreator{
		logger:            log.Default(),
		recipesCollection: recipesCollection,
	}
}

type recipesCreator struct {
	defaultOrdered    bool
	logger            *log.Logger
	recipesCollection *mongo.Collection
}

func (r recipesCreator) CreateRecipes(ctx context.Context, recipes ...*business.Recipe) error {
	documents := make([]any, 0, len(recipes))

	for _, recipe := range recipes {
		r.logger.Println("Recipe ID:", recipe.ID)
		documents = append(documents, NewRecipe(recipe))
	}

	opts := &options.InsertManyOptions{Ordered: &r.defaultOrdered}

	// TODO: use batches of {some limit} to insert many records
	// TODO: save recipe ingredients using a transaction
	_, err := r.recipesCollection.InsertMany(ctx, documents, opts)
	if err != nil {
		r.logger.Println(err)
		return err
	}

	return err
}

func NewRecipeCreatorForCLI(collectionFunc func() (*mongo.Collection, error)) business.RecipesCreatorFunc {
	return func(ctx context.Context, recipes ...*business.Recipe) error {
		recipesCollection, err := collectionFunc()
		if err != nil {
			return err
		}

		recipeCreator := NewRecipeCreator(recipesCollection)

		return recipeCreator.CreateRecipes(ctx, recipes...)
	}
}
