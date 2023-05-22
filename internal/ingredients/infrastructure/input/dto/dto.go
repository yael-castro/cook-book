package dto

import (
	"github.com/yael-castro/cb-search-engine-api/internal/ingredients/business/model/mass"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"
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
	*NutritionalInformation `json:"nutritional_information,omitempty"`
	ID                      int64  `json:"id"`
	Name                    string `json:"name,omitempty"`
	Description             string `json:"description,omitempty"`
}

type NutritionalInformation struct {
	Calories int64     `json:"calories,omitempty"`
	Fats     mass.Mass `json:"fats,omitempty"`
	Proteins mass.Mass `json:"proteins,omitempty"`
	Carbs    mass.Mass `json:"carbs,omitempty"`
	Fiber    mass.Mass `json:"fiber,omitempty"`
}
