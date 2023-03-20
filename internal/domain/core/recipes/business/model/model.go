package model

type (
	// Recipe kitchen recipe data
	Recipe struct {
		ID          int64
		Name        string
		Description string
		Ingredients []int64
	}
)
