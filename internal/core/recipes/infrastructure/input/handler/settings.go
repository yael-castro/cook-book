package handler

import (
	"github.com/yael-castro/cb-search-engine-api/internal/lib/code"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/server"
	"net/http"
)

func ErrorHandler() server.ErrorHandler {
	return server.NewErrorHandler(server.ErrorHandlerConfig{
		Prefix: "RCPS",
		Codes: map[int][]code.Code{},
		Logger: nil,
	})
}

func RouteMap(creator http.HandlerFunc) server.RouteMap {
	return server.RouteMap{
		"/v1/recipes": {
			http.MethodGet: creator,
		},
	}
}
