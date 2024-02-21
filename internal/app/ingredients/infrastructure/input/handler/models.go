package handler

import (
	"github.com/yael-castro/cook-book/internal/app/ingredients/business"
)

type Mass = business.Mass

func NewIngredient(ingredient business.Ingredient) Ingredient {
	return Ingredient{
		NutritionalInformation: NewNutritionalInformation(ingredient.NutritionalInformation),
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

func (i Ingredient) ToBusiness() business.Ingredient {
	return business.Ingredient{
		NutritionalInformation: i.NutritionalInformation.ToBusiness(),
		ID:                     i.ID,
		Name:                   i.Name,
		Description:            i.Description,
	}
}

func NewNutritionalInformation(information business.NutritionalInformation) NutritionalInformation {
	return (NutritionalInformation)(information)
}

type NutritionalInformation struct {
	Calories int64 `json:"calories,omitempty"`
	Fats     Mass  `json:"fats,omitempty"`
	Proteins Mass  `json:"proteins,omitempty"`
	Carbs    Mass  `json:"carbs,omitempty"`
	Fiber    Mass  `json:"fiber,omitempty"`
}

func (i NutritionalInformation) ToBusiness() business.NutritionalInformation {
	return (business.NutritionalInformation)(i)
}

type IngredientPage struct {
	Total       uint64       `json:"total"`
	Ingredients []Ingredient `json:"ingredients"`
}

func NewIngredients(ingredientsSlice []business.Ingredient) []Ingredient {
	ingredients := make([]Ingredient, 0, len(ingredientsSlice))

	for _, ingredient := range ingredientsSlice {
		ingredients = append(ingredients, NewIngredient(ingredient))
	}

	return ingredients
}

func BusinessIngredients(ingredientsSlice []Ingredient) []business.Ingredient {
	ingredients := make([]business.Ingredient, 0, len(ingredientsSlice))

	for _, ingredient := range ingredientsSlice {
		ingredients = append(ingredients, ingredient.ToBusiness())
	}

	return ingredients
}
