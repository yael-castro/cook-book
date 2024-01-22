package business

import (
	"context"
)

func NewRecipeCreator(saver RecipesSaver) RecipesCreator {
	if saver == nil {
		panic("missing dependencies")
	}

	return recipesCreator{
		saver: saver,
	}
}

type recipesCreator struct {
	saver RecipesSaver
}

func (r recipesCreator) CreateRecipes(ctx context.Context, recipes ...*Recipe) error {
	for _, recipe := range recipes {
		if err := recipe.Validate(); err != nil {
			return err
		}
	}

	return r.saver.SaveRecipes(ctx, recipes...)
}

func NewRecipeGenerator(writer RecipesWriter, saver RecipesSaver) RecipesGenerator {
	switch any(nil) {
	case writer, saver:
		panic("missing settings")
	}

	return &recipesGenerator{
		writer: writer,
		saver:  saver,
	}
}

type recipesGenerator struct {
	writer RecipesWriter
	saver  RecipesSaver
}

func (r recipesGenerator) GenerateRecipes(ctx context.Context, generate GenerateRecipes) error {
	if err := generate.Validate(); err != nil {
		return err
	}

	recipes, err := r.writer.WriteRecipes(ctx, generate)
	if err != nil {
		return err
	}

	return r.saver.SaveRecipes(ctx, recipes...)
}

// NewRecipesSearcher builds a materialization for the port.RecipesSearcher interface
func NewRecipesSearcher(finder RecipesFinder) RecipesSearcher {
	if finder == nil {
		panic("nil dependency")
	}

	return recipesSearcher{
		finder: finder,
	}
}

type recipesSearcher struct {
	finder RecipesFinder
}

func (r recipesSearcher) SearchRecipes(ctx context.Context, filter *RecipeFilter) ([]*Recipe, error) {
	if err := filter.Validate(); err != nil {
		return nil, err
	}

	return r.finder.FindRecipes(ctx, filter)
}
