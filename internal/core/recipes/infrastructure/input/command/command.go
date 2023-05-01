package command

import (
	"context"
	"flag"
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/cli"
	"log"
)

const usageTemplate = `%s %s is a tool to generate random recipes

Usage:

  %s %s [options]

Options:

`

// RecipeListGenerator defines an application adapter to interprets the command line parameters
type RecipeListGenerator interface {
	// GenerateRecipeList takes raw command line parameters to parse it and generate a recipe list in the storage
	GenerateRecipeList(context.Context, ...string) error
}

// NewRecipeListGenerator builds an instance of the unique implementation for the RecipeListGenerator interface
func NewRecipeListGenerator(generator port.RecipeSetGenerator) RecipeListGenerator {
	return &recipeListGenerator{
		RecipeSetGenerator: generator,
	}
}

type recipeListGenerator struct {
	RecipeSetGenerator port.RecipeSetGenerator
}

func (r recipeListGenerator) GenerateRecipeList(ctx context.Context, args ...string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	ingredients := flags.Uint64("ingredients", 0, "indicates the number of the ingredients for recipe")
	recipes := flags.Uint64("recipes", 0, "indicates the number of the recipes that will generated")

	err := flags.Parse(args[1:])
	if err != nil {
		return err
	}

	if *ingredients < 1 {
		r.usage(flags)
		return nil
	}

	if *recipes < 1 {
		r.usage(flags)
		return nil
	}

	return r.RecipeSetGenerator.GenerateRecipeSet(ctx, uint32(*recipes), uint32(*ingredients))
}

func (r recipeListGenerator) usage(set *flag.FlagSet) {
	log.Printf(
		usageTemplate,
		cli.BinaryName,
		set.Name(),
		cli.BinaryName,
		set.Name(),
	)
	set.PrintDefaults()
}
