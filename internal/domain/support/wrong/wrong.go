// Package wrong is a support subdomain that contains the errors related to the business
package wrong

var _ error = Validation("")

// Validation returns when occurs a validation error
type Validation string

// Error returns the string value of Validation
func (e Validation) Error() string {
	return string(e)
}

var _ error = NotFound("")

// NotFound error implementation
//
// Returns when occurs an error related to missing resource
type NotFound string

// Error returns the string value of NotFound
func (e NotFound) Error() string {
	return string(e)
}
