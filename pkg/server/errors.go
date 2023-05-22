package server

import (
	"errors"
	"github.com/yael-castro/cb-search-engine-api/pkg/errors/code"
	"github.com/yael-castro/cb-search-engine-api/pkg/server/response"
	"log"
	"net/http"
)

// Logger defines a server logger
type Logger interface {
	// Log format and prints input console received arguments
	Log(...any)
}

// ErrorConvertor defines a interface to convert errors input a http response
type ErrorConvertor interface {
	// ConvertError converts an error input a status code and also a object to response
	ConvertError(error) (int, *response.Common)
}

// ErrorHandler defines a error handler to manage the internal errors
type ErrorHandler interface {
	// HandleError handle err and parse it input a code.Code to then translate it input an integer status
	HandleError(http.ResponseWriter, *http.Request, error)
}

// ErrorHandlerConfig indicates how to build an instance of ErrorHandler
type ErrorHandlerConfig struct {
	Prefix string
	Codes  map[int][]code.Code
	Logger Logger
}

// NewErrorHandler builds an ErrorHandler based on a "map[int][]code.Code" which is used as configuration
//
// NOTES:
// - The error handler also implements the ErrorConverter interface
func NewErrorHandler(config ErrorHandlerConfig) ErrorHandler {
	handler := &errorHandler{
		prefix: config.Prefix,
		codes:  make(map[code.Code]*int),
		logger: config.Logger,
	}

	for status, codes := range config.Codes {
		handler.setStatus(status, codes...)
	}

	return handler
}

// _ "implement" constraint for the errorHandler struct
var _ ErrorConvertor = (*errorHandler)(nil)

type errorHandler struct {
	prefix string
	logger Logger
	codes  map[code.Code]*int
}

func (e *errorHandler) setStatus(status int, codes ...code.Code) {
	for _, v := range codes {
		e.codes[v] = &status
	}
}

func (e *errorHandler) Log(a ...any) {
	if e.logger != nil {
		e.logger.Log(a)
	}
}

func (e *errorHandler) ConvertError(err error) (int, *response.Common) {
	if err == nil {
		panic("nil error")
	}

	internal, ok := err.(code.Code)
	if !ok {
		internal, ok = errors.Unwrap(err).(code.Code)
		if !ok {
			e.Log(err)
			return http.StatusInternalServerError, &response.Common{
				Code:    e.prefix + code.Zero.Error(),
				Message: "an unexpected or unhandled error occurred",
			}
		}
	}

	status, ok := e.codes[internal]
	if !ok {
		e.Log("unknown internal status code: ", status)
		return http.StatusInternalServerError, &response.Common{
			Code:    e.prefix + internal.Error(),
			Message: err.Error(),
		}
	}

	return *status, &response.Common{
		Code:    e.prefix + internal.Error(),
		Message: err.Error(),
	}
}

func (e *errorHandler) HandleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("ERROR", err)
	status, common := e.ConvertError(err)
	_ = JSON(w, status, common)
}
