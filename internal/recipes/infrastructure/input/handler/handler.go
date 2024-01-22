package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"github.com/yael-castro/cb-search-engine-api/pkg/set"
	"net/http"
	"strconv"
	"strings"
)

func PostRecipes(creator business.RecipesCreator) echo.HandlerFunc {
	if creator == nil {
		panic("nil dependency")
	}

	return func(c echo.Context) error {
		recipes := make([]*Recipe, 0)

		// Unmarshal the request body
		if err := c.Bind(&recipes); err != nil {
			return err
		}

		// Conversion between data types
		arr := make([]*business.Recipe, 0, len(recipes))

		for _, recipe := range recipes {
			arr = append(arr, recipe.ToBModel())
		}

		// Creates many recipes
		err := creator.CreateRecipes(c.Request().Context(), arr...)
		if err != nil {
			return err
		}

		// Success response
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success operation",
		})
	}
}

// GetRecipes builds an instance of the unique implementation for the RecipeProvider interface based on a port.RecipesSearcher
func GetRecipes(searcher business.RecipesSearcher) echo.HandlerFunc {
	if searcher == nil {
		panic("nil dependency")
	}

	return func(c echo.Context) (err error) {
		q := c.QueryParams()

		filter := business.RecipeFilter{
			Ingredients: make(set.Set[int64]),
		}

		filter.Page, err = strconv.ParseUint(q.Get("page"), 10, 64)
		if err != nil {
			err = echo.NewHTTPError(http.StatusBadRequest, "missing query param 'page'")
			return
		}

		filter.Size, err = strconv.ParseUint(q.Get("size"), 10, 64)
		if err != nil {
			err = echo.NewHTTPError(http.StatusBadRequest, "missing query param 'size'")
			return
		}

		// Decoding ingredient values
		ingredient := int64(0)
		ingredients := strings.SplitN(q.Get("ingredients"), ",", 10)

		for _, v := range ingredients {
			ingredient, err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				err = fmt.Errorf("%w: ingredient id '%s' is not valid number", business.ErrInvalidIngredientID, v)
				return
			}

			filter.Ingredients.Add(ingredient)
		}

		// Searching recipes
		results, err := searcher.SearchRecipes(c.Request().Context(), &filter)
		if err != nil {
			return err
		}

		// Encoding results
		recipes := make([]*Recipe, 0, len(results))

		for i := range results {
			recipes = append(recipes, NewRecipe(results[i]))
		}

		return c.JSON(http.StatusOK, RecipePage{
			Recipes: recipes,
			Total:   filter.Total,
		})
	}
}
