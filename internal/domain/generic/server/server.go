// Package server contains everything related to api in a client-api model
//
// NOTE: this package should be in a library
package server

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server/response"
	"net/http"
	"strings"
)

type Configuration struct {
	RouteMap map[string]map[string]http.Handler
}

// Server defines a api in the model client-api
type Server interface {
	// Serve running the server on the address specified until the context is canceled
	Serve(context.Context, string) error
}

// New builds a http api that complies the Server interface based on the Router received
func New(config Configuration) Server {
	routes := config.RouteMap

	s := &server{
		ServeMux: http.NewServeMux(),
	}

	for route := range routes {
		for method, handler := range routes[route] {
			s.setRoute(method, route, handler)
		}
	}

	return s
}

type server struct {
	*http.ServeMux
}

func (s *server) setRoute(method string, path string, handler http.Handler) {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	s.ServeMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			JSON(w, http.StatusMethodNotAllowed, &response.Common{
				Message: fmt.Sprintf("method '%s' is not allowed", r.Method),
			})
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (s *server) Serve(ctx context.Context, address string) error {
	go http.ListenAndServe(address, s.ServeMux)

	<-ctx.Done()

	return ctx.Err()
}
