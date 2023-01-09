// Package server contains everything related to server
package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/wrong"
	"net/http"
)

type Configuration struct {
	SearchEngine echo.HandlerFunc
	ErrorHandler echo.HTTPErrorHandler
}

// New builds a http server that complies the Server interface based on the Configuration received
func New(config Configuration) Server {
	e := echo.New()

	if config.ErrorHandler != nil {
		e.HTTPErrorHandler = config.ErrorHandler
	}

	e.GET("/v1/recipes/suggestions", config.SearchEngine)

	return &server{Echo: e}
}

// Server defines a server in the model client-server
type Server interface {
	// Serve running the server on the address specified until the context is canceled
	Serve(context.Context, string) error
}

type server struct {
	*echo.Echo
}

func (s server) Serve(ctx context.Context, address string) error {
	go func() {
		<-ctx.Done()
		s.Close()
	}()

	return s.Start(address)
}

// DefaultErrorHandler default error handler to manage the http status code based on a error
func DefaultErrorHandler(err error, c echo.Context) {
	switch err := err.(type) {
	case wrong.Validation:
		c.JSON(http.StatusBadRequest, err.Error())
	case *echo.HTTPError:
		c.JSON(err.Code, err.Error())
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
