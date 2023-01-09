package handler

type (
	// recipePage is a data transfer object for the recipe page
	recipePage struct {
		Items      any    `json:"items"`
		TotalPages uint64 `json:"totalPages"`
		TotalItems uint64 `json:"totalItems"`
	}

	// recipeItem is a data transfer object for the recipe data
	recipeItem struct {
		Id          int64    `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Ingredients []uint32 `json:"ingredients"`
	}
)
