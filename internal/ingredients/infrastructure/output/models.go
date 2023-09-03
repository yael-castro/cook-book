package output

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business"
)

func BusinessIngredient(ingredient *Ingredient) *business.Ingredient {
	return &business.Ingredient{
		NutritionalInformation: (*business.NutritionalInformation)(ingredient.NutritionalInformation),
		ID:                     ingredient.ID,
		Name:                   ingredient.Name,
		Description:            ingredient.Description,
	}
}

func NewIngredient(ingredient *business.Ingredient) *Ingredient {
	return &Ingredient{
		NutritionalInformation: (*NutritionalInformation)(ingredient.NutritionalInformation),
		ID:                     ingredient.ID,
		Name:                   ingredient.Name,
		Description:            ingredient.Description,
	}
}

type Ingredient struct {
	*NutritionalInformation `bson:"nutritional_information,omitempty"`
	ID                      int64  `bson:"_id"`
	Name                    string `bson:",omitempty"`
	Description             string `bson:",omitempty"`
}

func (i Ingredient) IsValid() (bool, error) {
	return false, nil
}

type NutritionalInformation struct {
	Calories int64         `bson:",omitempty"`
	Fats     business.Mass `bson:",omitempty"`
	Proteins business.Mass `bson:",omitempty"`
	Carbs    business.Mass `bson:",omitempty"`
	Fiber    business.Mass `bson:",omitempty"`
}
