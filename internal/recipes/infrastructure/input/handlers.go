package input

import (
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/pkg/pagination"
	"github.com/yael-castro/cb-search-engine-api/pkg/server"
	"github.com/yael-castro/cb-search-engine-api/pkg/server/response"
	"net/http"
)

func NewRecipesCreator(adder business.RecipesAdder, handler server.ErrorHandler) http.HandlerFunc {
	switch any(nil) {
	case adder, handler:
		panic("nil dependency")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		recipes := make([]*Recipe, 0)

		// Unmarshal the request body
		if !server.Bind(w, r, &recipes) {
			return
		}

		// Conversion between data types
		arr := make([]*business.Recipe, 0, len(recipes))

		for _, recipe := range recipes {
			arr = append(arr, BusinessRecipe(recipe))
		}

		// Creates many recipes
		err := adder.AddRecipes(r.Context(), arr...)
		if err != nil {
			handler.HandleError(w, r, err)
			return
		}

		// Success response
		server.JSON(w, http.StatusCreated, response.Common{
			Message: "success operation",
		})
	}
}

// NewRecipesFinder builds an instance of the unique implementation for the RecipeProvider interface based on a port.RecipesSearcher
func NewRecipesFinder(finder business.RecipesFinder, handler server.ErrorHandler) http.HandlerFunc {
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

		items := make([]*Recipe, 0, len(recipes))
		for _, recipe := range recipes {
			items = append(items, NewRecipe(recipe))
		}

		page := RecipePage{
			Items:      items,
			TotalPages: p.Pages(),
			TotalItems: p.TotalResults(),
		}

		_ = server.JSON(w, http.StatusOK, page)
	}
}
