package output

import (
	"context"
	ingredients "github.com/yael-castro/cb-search-engine-api/internal/ingredients/infrastructure/output"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// NewRecipeCreator builds an adapter for the port.RecipesCreator
func NewRecipeCreator(database *mongo.Database, logger *log.Logger) business.RecipesCreator {
	switch any(nil) {
	case logger, database:
		panic("missing settings")
	}

	return &recipesCreator{
		logger:                logger,
		client:                database.Client(),
		recipesCollection:     database.Collection("recipes"),
		ingredientsCollection: database.Collection("ingredients"),
	}
}

type recipesCreator struct {
	false                 bool
	logger                *log.Logger
	client                *mongo.Client
	recipesCollection     *mongo.Collection
	ingredientsCollection *mongo.Collection
}

func (r recipesCreator) CreateRecipes(ctx context.Context, recipes ...*business.Recipe) error {
	recipeDocuments := make([]any, 0, len(recipes))
	ingredientDocuments := make([]any, 0, len(recipes))

	for _, recipe := range recipes {
		recipeDocuments = append(recipeDocuments, NewRecipe(recipe))

		for _, ingredient := range recipe.Ingredients {
			ingredientDocuments = append(ingredientDocuments, ingredients.NewIngredient(ingredient))
		}
	}

	// TODO: avoid reflection to serialize the recipe and ingredient objects
	transaction := func(ctx mongo.SessionContext) (any, error) {
		opts := &options.InsertManyOptions{Ordered: &r.false}

		_, err := r.recipesCollection.InsertMany(ctx, recipeDocuments, opts)
		if err != nil {
			r.logger.Println(err)
			return nil, err
		}

		_, err = r.ingredientsCollection.InsertMany(ctx, ingredientDocuments, opts)
		return nil, err
	}

	session, err := r.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, transaction)
	return err
}

func NewRecipeCreatorFunc(databaseFunc func() (*mongo.Database, error), logger *log.Logger) business.RecipesCreatorFunc {
	return func(ctx context.Context, recipes ...*business.Recipe) error {
		database, err := databaseFunc()
		if err != nil {
			return err
		}

		recipeCreator := NewRecipeCreator(database, logger)

		return recipeCreator.CreateRecipes(ctx, recipes...)
	}
}
