package business

import (
	"github.com/yael-castro/cb-search-engine-api/pkg/errors/code"
)

const (
	_ code.Code = iota
	InvalidMin
	InvalidMax
	// MissingIngredientIdentifiers indicates that the ingredients identifiers are not present and are required
	// to perform an action
	MissingIngredientIdentifiers
	// InvalidIngredientID indicates that the ingredient id not valid
	InvalidIngredientID
)

const (
	minIngredientNumber = 1
	maxIngredientNumber = 30
)
