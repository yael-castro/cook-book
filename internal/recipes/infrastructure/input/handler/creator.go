package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/infrastructure/input/dto"
	"github.com/yael-castro/cb-search-engine-api/pkg/server"
	"github.com/yael-castro/cb-search-engine-api/pkg/server/response"
	"net/http"
)

func NewRecipeCreator(manager port.RecipeManager, handler server.ErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipes := make([]*dto.Recipe, 0)

		// Unmarshal the request body
		if !server.Bind(w, r, &recipes) {
			return
		}

		// Conversion between data types
		arr := make([]*model.Recipe, 0, len(recipes))

		for _, recipe := range recipes {
			arr = append(arr, dto.ToModelRecipe(recipe))
		}

		// Creates many recipes
		err := manager.CreateRecipes(r.Context(), arr...)
		if err != nil {
			handler.HandleError(w, r, err)
			return
		}

		// Success response
		server.JSON(w, http.StatusCreated, response.Common{
			Message: "Success operation",
		})
	}
}
