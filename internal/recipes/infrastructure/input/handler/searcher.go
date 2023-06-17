package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input/dto"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
	"github.com/yael-castro/cb-search-engine-api/pkg/server"
	"net/http"
)

// NewRecipesFinder builds an instance of the unique implementation for the RecipeProvider interface based on a port.RecipesSearcher
func NewRecipesFinder(finder port.RecipesFinder, handler server.ErrorHandler) http.HandlerFunc {
	switch any(nil) {
	case finder, handler:
		panic("nil dependency")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		p := pagination.New(q.Get("page"), q.Get("limit"))

		recipes, err := finder.FindRecipes(r.Context(), q.Get("ingredients"), p)
		if err != nil {
			handler.HandleError(w, r, err)
			return
		}

		items := make([]*dto.Recipe, 0, len(recipes))
		for _, recipe := range recipes {
			items = append(items, dto.NewRecipe(recipe))
		}

		page := dto.RecipePage{
			Items:      items,
			TotalPages: p.Pages(),
			TotalItems: p.TotalResults(),
		}

		_ = server.JSON(w, http.StatusOK, page)
	}
}
