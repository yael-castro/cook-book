// Package server contains everything related to api in a client-api model
package server

import (
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/server/response"
	"net/http"
	"strings"
)

type RouteMap map[string]map[string]http.Handler

type Config struct {
	Maps []RouteMap
}

// New builds a http api that complies the Server interface based on the Router received
func New(config Config) http.Handler {
	s := &server{}

	for _, mapper := range config.Maps {
		for route, routes := range mapper {
			for method, handler := range routes {
				s.setRoute(method, route, handler)
			}
		}
	}

	return s
}

type server map[string]map[string]http.Handler

func (s *server) setRoute(method string, path string, handler http.Handler) {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	if _, exists := (*s)[path]; !exists {
		(*s)[path] = make(map[string]http.Handler)
	}

	(*s)[path][method] = handler
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

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
