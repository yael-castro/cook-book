package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/domain/core/finder/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/code"
	"net/http"
)

func NewErrorHandler() server.ErrorHandler {
	return server.NewErrorHandler(server.ErrorHandlerConfig{
		Prefix: "SeEn",
		Codes: map[int][]code.Code{
			http.StatusBadRequest: {
				consts.InvalidIngredientID,
				consts.MissingIngredientIdentifiers,
			},
		},
	})
}
