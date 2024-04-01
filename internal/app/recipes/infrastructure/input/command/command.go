package command

import (
	"github.com/spf13/cobra"
	"github.com/yael-castro/cook-book/internal/app/recipes/business"
)

// Flag names
const (
	recipesFlag     = "recipes"
	ingredientsFlag = "ingredients"
)

// GenerateRecipes builds an instance of the unique implementation for the RecipeListGenerator interface
func GenerateRecipes(generator business.RecipesGenerator) *cobra.Command {
	cmd := &cobra.Command{
		Use:       "generate",
		Short:     "generates a set of recipes based on different algorithms",
		ValidArgs: []string{ingredientsFlag, recipesFlag},
		Run:       generateRecipesFunc(generator),
	}

	flags := cmd.PersistentFlags()

	_ = flags.Uint32(ingredientsFlag, 0, "indicates the number of the ingredients for recipe")
	_ = flags.Uint32(recipesFlag, 0, "indicates the number of the recipes that will generated")

	_ = cmd.MarkPersistentFlagRequired(ingredientsFlag)
	_ = cmd.MarkPersistentFlagRequired(recipesFlag)

	return cmd
}

func generateRecipesFunc(generator business.RecipesGenerator) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := cmd.ParseFlags(args); err != nil {
			return
		}

		recipes, err := cmd.Flags().GetUint32(recipesFlag)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		ingredients, err := cmd.Flags().GetUint32(ingredientsFlag)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		generate := business.GenerateRecipes{
			Ingredients: ingredients,
			Recipes:     recipes,
		}

		cmd.Printf(
			"Generating %d recipes with %d ingredients...\n",
			generate.Recipes,
			generate.Ingredients,
		)

		if err := generator.GenerateRecipes(cmd.Context(), generate); err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Println("Success!")
	}
}
