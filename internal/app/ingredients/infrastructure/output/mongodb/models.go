package mongodb

import (
	"github.com/yael-castro/cook-book/internal/app/ingredients/business"
	"go.mongodb.org/mongo-driver/bson"
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

func (i Ingredient) ToBusinessModel() business.Ingredient {
	return business.Ingredient{
		NutritionalInformation: i.NutritionalInformation.ToBusiness(),
		ID:                     i.ID,
		Name:                   i.Name,
		Description:            i.Description,
	}
}

type Ingredients []Ingredient

func (i Ingredients) ToBusinessModel() business.Ingredients {
	ingredients := make([]business.Ingredient, len(i))

	for index := range i {
		ingredients[index] = i[index].ToBusinessModel()
	}

	return ingredients
}

func NewIngredientFilter(ingredientFilter business.IngredientFilter) IngredientFilter {
	return (IngredientFilter)(ingredientFilter)
}

// IngredientFilter contains parameters to make ingredient searches
type IngredientFilter struct {
	Page    uint64 `bson:"-"`
	Size    uint64 `bson:"-"`
	Total   uint64 `bson:"-"`
	Keyword string `bson:"-"`
	Random  bool   `bson:"-"`
}

func (i IngredientFilter) Pipeline() bson.A {
	return bson.A{
		bson.D{
			{
				Key: "$sample",
				Value: bson.D{
					{
						Key:   "size",
						Value: i.Size,
					},
				},
			},
		},
	}
}

func (i IngredientFilter) Document() bson.D {
	return bson.D{
		{
			Key: "name",
			Value: bson.D{
				{
					Key:   "$regex",
					Value: i.Keyword,
				},
			},
		},
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
