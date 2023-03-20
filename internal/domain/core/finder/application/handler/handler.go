package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/application/dto"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/pagination"
	"net/http"
)

// NewRecipeEngine builds an instance of the unique implementation for the RecipeProvider interface based on a port.RecipeSearcher
func NewRecipeEngine(searcher port.RecipeSearcher, handler server.ErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		p := pagination.New(q.Get("page"), q.Get("limit"))

		recipes, err := searcher.SearchRecipe(r.Context(), q.Get("ingredients"), p)
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
