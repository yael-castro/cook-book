package business

import (
	"errors"
)

type Ingredient struct {
	*NutritionalInformation
	ID                int64
	Name, Description string
}

func (i *Ingredient) Validate() error {
	if i == nil {
		return errors.New("missing ingredient data")
	}

	if err := i.NutritionalInformation.Validate(); err != nil {
		return err
	}

	switch {
	case i.ID < 1:
		return errors.New("missing ingredient id")
	case i.Name == "":
		return errors.New("missing ingredient name")
	case i.Description == "":
		return errors.New("missing ingredient description")
	}

	return nil
}

type NutritionalInformation struct {
	Calories                     int64
	Fats, Proteins, Carbs, Fiber Mass
}

func (i *NutritionalInformation) Validate() error {
	if i == nil {
		return errors.New("missing nutritional information")
	}

	return nil
}
