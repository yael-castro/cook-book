package input

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
	"net/http"
)

func ErrorHandler(handler echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		as := business.RecipeError(0)

		if errors.As(err, &as) {
			message := echo.Map{
				"code":    as.Error(),
				"message": err.Error(),
			}

			switch as {
			case
				business.ErrInvalidRecipe,
				business.ErrInvalidPageSize,
				business.ErrInvalidIngredients,
				business.ErrInvalidIngredientID:
				_ = c.JSON(http.StatusBadRequest, message)
			}

			return
		}

		handler(err, c)
	}
}
