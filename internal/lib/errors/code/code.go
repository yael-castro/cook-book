// Package code contains internal status codes that are translate in codes
package code

import (
	"errors"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/errors/wrapper"
	"strconv"
)

// Zero indicates a unidentified Code
const Zero Code = 0

// New builds a Code wrap with an embed message
func New(code Code, message string) error {
	return wrapper.New(code, errors.New(message))
}

var _ error = Code(0)

// Code customized error codes to better error handling
type Code int64

// Error returns the value of Code with the prefix
func (c Code) Error() string {
	return strconv.FormatInt(int64(c), 10)
}
