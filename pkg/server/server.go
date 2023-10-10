// Package server contains everything related to api input a client-api model
package server

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/pkg/server/response"
	"net/http"
)

// RouteMap defines the ways that http requests can follow depending on the path and method
type RouteMap map[string]map[string]http.Handler

// New builds an instance of http.Handler based on the instances of RouteMap which defines how the requests should be managed
func New(routeMaps ...RouteMap) http.Handler {
	s := &server{}

	for _, routeMap := range routeMaps {
		for route, routes := range routeMap {
			for method, handler := range routes {
				s.setRoute(method, route, handler)
			}
		}
	}

	return s
}

type server map[string]map[string]http.Handler

func (s *server) setRoute(method string, path string, handler http.Handler) {
	for path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	if _, exists := (*s)[path]; !exists {
		(*s)[path] = make(map[string]http.Handler)
	}

	(*s)[path][method] = handler
}

// ServeHTTP handles a http request and handle it depending on the request path and method
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // TODO: evaluate if the pool pattern is required

	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	// Validates if the path exists
	if (*s)[r.URL.Path] == nil {
		JSON(w, http.StatusNotFound, &response.Common{
			Message: fmt.Sprintf("the path '%s' does not exists", path),
		})
		return
	}

	// Validates if the path allows the method
	if (*s)[path][r.Method] == nil {
		JSON(w, http.StatusMethodNotAllowed, &response.Common{
			Message: fmt.Sprintf("method '%s' is not allowed", r.Method),
		})
		return
	}

	(*s)[path][r.Method].ServeHTTP(w, r)
}

// String builds and returns a representative for the server instance
func (s *server) String() string {
	str := ""

	for route := range *s {
		for method := range (*s)[route] {
			str += fmt.Sprintf("%s %s - %T\n", route, method, (*s)[route][method])
		}
	}

	return str
}
