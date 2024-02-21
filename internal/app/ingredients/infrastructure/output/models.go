package output

import (
	"github.com/yael-castro/cook-book/internal/app/ingredients/business"
)

func NewIngredient(ingredient business.Ingredient) Ingredient {
	return Ingredient{
		NutritionalInformation: (NutritionalInformation)(ingredient.NutritionalInformation),
		ID:                     ingredient.ID,
		Name:                   ingredient.Name,
		Description:            ingredient.Description,
	}
}

type Ingredient struct {
	NutritionalInformation `bson:"nutritional_information,omitempty"`
	ID                     int64  `bson:"_id"`
	Name                   string `bson:",omitempty"`
	Description            string `bson:",omitempty"`
}

// ToBusiness transforms the Ingredient object into an object that the business can understand.
func (i Ingredient) ToBusiness() business.Ingredient {
	return business.Ingredient{
		NutritionalInformation: i.NutritionalInformation.ToBusiness(),
		ID:                     i.ID,
		Name:                   i.Name,
		Description:            i.Description,
	}
}

type NutritionalInformation struct {
	Calories int64         `bson:",omitempty"`
	Fats     business.Mass `bson:",omitempty"`
	Proteins business.Mass `bson:",omitempty"`
	Carbs    business.Mass `bson:",omitempty"`
	Fiber    business.Mass `bson:",omitempty"`
}

// ToBusiness transforms the NutritionalInformation object into an object that the business can understand.
func (n NutritionalInformation) ToBusiness() business.NutritionalInformation {
	return (business.NutritionalInformation)(n)
}
