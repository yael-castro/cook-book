package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business"
	"net/http"
	"strconv"
)

func GetIngredients(searcher business.IngredientSearcher) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := c.QueryParams()
		ctx := c.Request().Context()

		filter := business.IngredientFilter{}

		// Decoding query params
		filter.Keyword = q.Get("q")
		filter.Page, _ = strconv.ParseUint(q.Get("page"), 10, 64)
		filter.Size, _ = strconv.ParseUint(q.Get("size"), 10, 64)
		filter.Random, _ = strconv.ParseBool(q.Get("random"))

		ingredients, err := searcher.SearchIngredients(ctx, &filter)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, IngredientPage{
			Total:       filter.Total,
			Ingredients: NewIngredients(ingredients),
		})
	}
}
