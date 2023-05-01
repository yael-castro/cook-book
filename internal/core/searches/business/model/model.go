package model

import (
	"bytes"
	"fmt"
	recipes "github.com/yael-castro/cb-search-engine-api/internal/core/recipes/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/pagination"
	"github.com/yael-castro/cb-search-engine-api/internal/lib/set"
	"golang.org/x/crypto/sha3"
	"sort"
	"strconv"
)

var _ fmt.Stringer = (*RecipeFilter)(nil)

// RecipeFilter filter for paginated recipe searches
type RecipeFilter struct {
	*pagination.Pagination
	// Ingredients are the ingredients used to perform a recipe search
	Ingredients set.Set[int64]
}

// String returns a sha-3 hash of 28 bytes that works a token to identify the object.
// The token is build based comma-separated values that contains the ingredients identifiers and the paginate options.
//
// Example:
//
// SHA-3(1, 2, 3, 100, 20)
func (f RecipeFilter) String() string {
	// No action is required
	if len(f.Ingredients) < 1 {
		return ""
	}

	// Builds a buffer
	buffer := bytes.NewBuffer(nil)

	// Sorts the slice
	slice := f.Ingredients.Slice()

	sort.Slice(slice, func(x, y int) bool {
		return slice[x] > slice[y]
	})

	// Writes in the buffer the ingredient identifiers separated by commas
	for _, v := range slice {
		buffer.WriteString(strconv.FormatInt(v, 10))
		buffer.WriteRune(',')
	}

	// Writes in the buffer the pagination values separated by commas
	buffer.WriteString(strconv.FormatInt(int64(f.Limit()), 10))

	buffer.WriteRune(',')

	buffer.WriteString(strconv.FormatInt(int64(f.Page()), 10))

	// Builds a hash
	hash := sha3.Sum224(buffer.Bytes())
	return string(hash[:])
}

// Recipe alias for recipes.Recipe
type Recipe = recipes.Recipe
