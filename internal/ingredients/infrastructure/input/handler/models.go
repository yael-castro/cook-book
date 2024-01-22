package handler

import (
	ingredients "github.com/yael-castro/cb-search-engine-api/internal/ingredients/business"
	"github.com/yael-castro/cb-search-engine-api/internal/recipes/business"
)

func BusinessIngredients(ingredientsSlice []Ingredient) []business.Ingredient {
	ingredients := make([]business.Ingredient, 0, len(ingredientsSlice))

	for _, ingredient := range ingredientsSlice {
		ingredients = append(ingredients, ingredient.ToBModel())
	}

	return ingredients
}

func NewIngredients(ingredientsSlice []business.Ingredient) []Ingredient {
	ingredients := make([]Ingredient, 0, len(ingredientsSlice))

	for _, ingredient := range ingredientsSlice {
		ingredients = append(ingredients, NewIngredient(ingredient))
	}

	return ingredients
}

func NewIngredient(ingredient business.Ingredient) Ingredient {
	return Ingredient{
		NutritionalInformation: (NutritionalInformation)(ingredient.NutritionalInformation),
		ID:                     ingredient.ID,
		Name:                   ingredient.Name,
		Description:            ingredient.Description,
	}
}

type Ingredient struct {
	NutritionalInformation `json:"nutritional_information,omitempty"`
	ID                     int64  `json:"id"`
	Name                   string `json:"name,omitempty"`
	Description            string `json:"description,omitempty"`
}

func (i Ingredient) ToBModel() business.Ingredient {
	return business.Ingredient{
		NutritionalInformation: (business.NutritionalInformation)(i.NutritionalInformation),
		ID:                     i.ID,
		Name:                   i.Name,
		Description:            i.Description,
	}
}

type NutritionalInformation struct {
	Calories int64 `json:"calories,omitempty"`
	Fats     Mass  `json:"fats,omitempty"`
	Proteins Mass  `json:"proteins,omitempty"`
	Carbs    Mass  `json:"carbs,omitempty"`
	Fiber    Mass  `json:"fiber,omitempty"`
}

type Mass = ingredients.Mass
