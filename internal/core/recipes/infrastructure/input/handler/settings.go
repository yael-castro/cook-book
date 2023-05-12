package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/model/consts"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/errors/code"
	server2 "github.com/yael-castro/cb-search-engine-api/internal/lib/server"
	"net/http"
)

func ErrorHandler() server2.ErrorHandler {
	return server2.NewErrorHandler(server2.ErrorHandlerConfig{
		Prefix: "RCPS",
		Codes: map[int][]code.Code{
			http.StatusBadRequest: {
				consts.InvalidMin,
				consts.InvalidMax,
			},
		},
		Logger: nil,
	})
}

func RouteMap(creator http.HandlerFunc) server2.RouteMap {
	return server2.RouteMap{
		"/v1/recipes": {
			http.MethodGet: creator,
		},
	}
}
