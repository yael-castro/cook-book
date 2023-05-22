package dto

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business/model"
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business/model/mass"
)

func NewIngredient(ingredient *model.Ingredient) *Ingredient {
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
	Calories int64     `bson:",omitempty"`
	Fats     mass.Mass `bson:",omitempty"`
	Proteins mass.Mass `bson:",omitempty"`
	Carbs    mass.Mass `bson:",omitempty"`
	Fiber    mass.Mass `bson:",omitempty"`
}
