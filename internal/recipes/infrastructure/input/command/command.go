package command

import (
	"context"
	"flag"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/pkg/cli"
	"log"
)

// NewRecipeListGenerator builds an instance of the unique implementation for the RecipeListGenerator interface
func NewRecipeListGenerator(generator port.RecipeSetGenerator) cli.Commander {
	return &recipeListGenerator{
		RecipeSetGenerator: generator,
	}
}

type recipeListGenerator struct {
	flags              *flag.FlagSet
	RecipeSetGenerator port.RecipeSetGenerator
}

// Command generates a recipe list based on the flags ingredients and recipes
func (r *recipeListGenerator) Command(ctx context.Context, args ...string) error {
	r.flags = flag.NewFlagSet(args[0], flag.ContinueOnError)

	ingredients := r.flags.Uint64("ingredients", 0, "indicates the number of the ingredients for recipe")
	recipes := r.flags.Uint64("recipes", 0, "indicates the number of the recipes that will generated")

	err := r.flags.Parse(args[1:])
	if err != nil {
		log.Println(err)
		return err
	}

	return r.RecipeSetGenerator.GenerateRecipeSet(ctx, uint32(*recipes), uint32(*ingredients))
}

// Help shows the instructions to use the Command
func (r *recipeListGenerator) Help() {
	cli.Usage(r.flags)
}
