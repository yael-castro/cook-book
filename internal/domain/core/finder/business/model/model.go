package model

import (
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/pagination"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/support/set"
)

// RecipeFilter filter for paginated recipe searches
type RecipeFilter struct {
	*pagination.Pagination
	// Ingredients are the ingredients used to perform a recipe search
	Ingredients set.Set[int32]
}
