package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/cook-book/internal/app/recipes/business"
	"net/http"
)

func ErrorHandler(handler echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		as := business.RecipeError(0)

		if !errors.As(err, &as) {
			handler(err, c)
			return
		}

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
		default:
			_ = c.JSON(http.StatusInternalServerError, message)
		}
	}
}
