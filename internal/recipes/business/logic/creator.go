package logic

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
)

func NewRecipeAdder(creator port.RecipesCreator) port.RecipesAdder {
	if creator == nil {
		panic("nil dependency")
	}

	return &recipeAdder{
		creator,
	}
}

type recipeAdder struct {
	port.RecipesCreator
}

func (r recipeAdder) AddRecipes(ctx context.Context, recipes ...*model.Recipe) error {
	for _, recipe := range recipes {
		if isValid, err := recipe.IsValid(); !isValid {
			return err
		}
	}

	return r.CreateRecipes(ctx, recipes...)
}
