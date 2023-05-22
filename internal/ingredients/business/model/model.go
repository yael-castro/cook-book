package model

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business/model/mass"
)

type Ingredient struct {
	*NutritionalInformation
	ID                int64
	Name, Description string
}

func (i Ingredient) IsValid() (bool, error) {
	return false, nil
}

type NutritionalInformation struct {
	Calories                     int64
	Fats, Proteins, Carbs, Fiber mass.Mass
}
