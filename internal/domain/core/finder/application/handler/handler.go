package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/port"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/pagination"
	"net/http"
)

// RecipeProvider defines a primary adapter to handle all http request related to recipe finder
type RecipeProvider interface {
	// ProvideRecipe handle http request to find recipes by ingredients
	ProvideRecipe(echo.Context) error
}

// NewRecipeProvider builds an instance of the unique implementation for the RecipeProvider interface based on a port.RecipeSearcher
func NewRecipeProvider(searcher port.RecipeSearcher) RecipeProvider {
	return recipeFinder{Searcher: searcher}
}

type recipeFinder struct {
	Searcher port.RecipeSearcher
}

func (r recipeFinder) ProvideRecipe(c echo.Context) error {
	p := pagination.New(c.QueryParam("page"), c.QueryParam("limit"))

	recipes, err := r.Searcher.SearchRecipe(c.Request().Context(), c.QueryParam("ingredients"), p)
	if err != nil {
		return err
	}

	items := make([]*recipeItem, 0, len(recipes))
	for _, recipe := range recipes {
		items = append(items, (*recipeItem)(recipe))
	}

	page := recipePage{
		Items:      items,
		TotalPages: p.Pages(),
		TotalItems: p.TotalResults(),
	}

	return c.JSON(http.StatusOK, page)
}
