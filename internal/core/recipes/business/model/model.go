package model

// Recipe kitchen recipe data
type Recipe struct {
	ID          int64
	Name        string
	Description string
	Ingredients []int64
}

func (r Recipe) IsValid() (bool, error) {
	return false, nil
}
