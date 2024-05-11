//go:build cli

package container

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/yael-castro/cook-book/internal/app/recipes/business"
	"github.com/yael-castro/cook-book/internal/app/recipes/infrastructure/input/command"
	"github.com/yael-castro/cook-book/internal/app/recipes/infrastructure/output/mock"
	"github.com/yael-castro/cook-book/internal/app/recipes/infrastructure/output/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

func Inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case *cobra.Command:
		return injectCommand(ctx, a)
	case *func() (*mongo.Database, error):
		return injectDatabaseFunc(ctx, a)
	}

	return inject(ctx, a)
}

func injectCommand(ctx context.Context, cmd *cobra.Command) (err error) {
	logger := log.Default()

	var databaseFunc func() (*mongo.Database, error)

	if err := Inject(ctx, &databaseFunc); err != nil {
		return err
	}

	// Driven adapters
	recipesWriter := mock.NewRecipesWriter()
	recipesSaverFunc := mongodb.NewRecipeSaverFunc(databaseFunc, logger)

	// Ports for driving adapters
	recipesGenerator := business.NewRecipeGenerator(recipesWriter, recipesSaverFunc)

	// Setting command
	*cmd = cobra.Command{
		Use:     os.Args[0],
		Short:   "Tool for managing the recipes operations.",
		Version: GitCommit,
	}

	cmd.CompletionOptions.DisableDefaultCmd = true

	// Setting drive adapters (sub-commands)
	cmd.AddCommand(command.GenerateRecipes(recipesGenerator))

	return
}

func injectDatabaseFunc(ctx context.Context, databaseFunc *func() (*mongo.Database, error)) error {
	*databaseFunc = func() (*mongo.Database, error) {
		var mongoDB mongo.Database

		if err := Inject(ctx, &mongoDB); err != nil {
			return nil, err
		}

		return &mongoDB, nil
	}
	return nil
}
