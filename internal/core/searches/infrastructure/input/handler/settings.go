package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/core/searches/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/code"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/server"
	"net/http"
)

func ErrorHandler() server.ErrorHandler {
	return server.NewErrorHandler(server.ErrorHandlerConfig{
		Prefix: "SRCH",
		Codes: map[int][]code.Code{
			http.StatusBadRequest: {
				consts.InvalidIngredientID,
				consts.MissingIngredientIdentifiers,
			},
		},
	})
}

func RouteMap(searcher http.Handler) server.RouteMap {
	return server.RouteMap{
		"/v1/recipes": {
			http.MethodGet: searcher,
		},
	}
}
