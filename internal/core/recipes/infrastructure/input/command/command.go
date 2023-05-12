package command

import (
	"context"
	"errors"
	"flag"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/cli"
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
func (r recipeListGenerator) Command(ctx context.Context, args ...string) error {
	r.flags = flag.NewFlagSet(args[0], flag.ExitOnError)

	ingredients := r.flags.Uint64("ingredients", 0, "indicates the number of the ingredients for recipe")
	recipes := r.flags.Uint64("recipes", 0, "indicates the number of the recipes that will generated")

	err := r.flags.Parse(args[1:])
	if err != nil {
		return err
	}

	if *ingredients < 1 {
		return errors.New("the number of ingredients are minor than the allowed ingredients")
	}

	if *recipes < 1 {
		return errors.New("the number of ingredients are mayor than allowed ingredients")
	}

	return r.RecipeSetGenerator.GenerateRecipeSet(ctx, uint32(*recipes), uint32(*ingredients))
}

// Help shows the instructions to use the Command
func (r recipeListGenerator) Help() {
	if r.flags == nil {
		return
	}

	cli.Usage(r.flags)
}