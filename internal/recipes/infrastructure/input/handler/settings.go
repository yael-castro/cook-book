package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/pkg/errors/code"
	"github.com/yael-castro/cb-search-engine-api/pkg/server"
	"net/http"
)

func ErrorHandler() server.ErrorHandler {
	return server.NewErrorHandler(server.ErrorHandlerConfig{
		Prefix: "RCPS",
		Codes: map[int][]code.Code{
			http.StatusBadRequest: {
				consts.InvalidMin,
				consts.InvalidMax,
				consts.InvalidIngredientID,
				consts.MissingIngredientIdentifiers,
			},
		},
		Logger: nil,
	})
}

func RouteMap(creator http.HandlerFunc, searcher http.HandlerFunc) server.RouteMap {
	return server.RouteMap{
		"/v1/recipes": {
			http.MethodGet:  searcher,
			http.MethodPost: creator,
		},
	}
}
