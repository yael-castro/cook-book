package writes

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/output/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// NewRecipeCreator builds an adapter for the port.RecipesCreator
func NewRecipeCreator(collection *mongo.Collection) port.RecipesCreator {
	if collection == nil {
		panic("missing MongoDB collection")
	}

	return &recipesCreator{
		logger:            log.Default(),
		recipesCollection: collection,
	}
}

type recipesCreator struct {
	defaultOrdered    bool
	logger            *log.Logger
	recipesCollection *mongo.Collection
}

func (r recipesCreator) CreateRecipes(ctx context.Context, recipes ...*model.Recipe) error {
	documents := make([]any, 0, len(recipes))

	for _, recipe := range recipes {
		log.Println("Recipe ID:", recipe.ID)
		documents = append(documents, dto.NewRecipe(recipe))
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
