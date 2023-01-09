package model

type (
	// Recipe kitchen recipe data
	Recipe struct {
		Id          int64 `json:"id" bson:"_id"`
		Name        string
		Description string
		Ingredients []uint32
	}
)
