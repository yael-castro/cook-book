package dto

import "github.com/yael-castro/cb-search-engine-api/internal/recipes/business/model"

func ToModelIngredient(ingredient *Ingredient) *model.Ingredient {
	return &model.Ingredient{
		NutritionalInformation: (*model.NutritionalInformation)(ingredient.NutritionalInformation),
		ID:                     ingredient.ID,
		Name:                   ingredient.Name,
		Description:            ingredient.Description,
	}
}
