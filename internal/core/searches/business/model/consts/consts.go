package consts

import "github.com/yael-castro/cb-search-engine-api/internal/lib/errors/code"

const (
	UnknownError code.Code = iota
	// MissingIngredientIdentifiers indicates that the ingredients identifiers are not present and are required
	// to perform an action
	MissingIngredientIdentifiers
	// InvalidIngredientID indicates that the ingredient id not valid
	InvalidIngredientID
)
