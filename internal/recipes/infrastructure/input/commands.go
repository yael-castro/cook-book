package input

import (
	"context"
	"flag"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/pkg/command"
)

// CommandGenerateRecipes builds an instance of the unique implementation for the RecipeListGenerator interface
func CommandGenerateRecipes(generator business.RecipesGenerator) command.Command {
	const commandName = "generate"

	return recipeListGenerator{
		flags:            flag.NewFlagSet(commandName, flag.ContinueOnError),
		recipesGenerator: generator,
	}
}

type recipeListGenerator struct {
	flags            *flag.FlagSet
	recipesGenerator business.RecipesGenerator
}

func (r recipeListGenerator) Name() string {
	return r.flags.Name()
}

func (r recipeListGenerator) Description() string {
	return "generates a set of random recipes"
}

func (r recipeListGenerator) Execute(ctx context.Context, args ...string) error {
	ingredients := r.flags.Uint64("ingredients", 0, "indicates the number of the ingredients for recipe")
	recipes := r.flags.Uint64("recipes", 0, "indicates the number of the recipes that will generated")

	err := r.flags.Parse(args[1:])
	if err != nil {
		return err
	}

	fmt.Println("Generating recipes...")

	if err := r.recipesGenerator.GenerateRecipes(ctx, uint32(*recipes), uint32(*ingredients)); err != nil {
		return err
	}

	fmt.Println("Successfully!")
	return nil
}
