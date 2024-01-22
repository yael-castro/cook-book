package business

import (
	"fmt"
)

const (
	G   Mass = 1
	DAG      = G * 10
	HG       = DAG * 10
	KG       = HG * 10
)

type Mass int64

type Ingredient struct {
	NutritionalInformation
	ID                int64
	Name, Description string
}

func (i Ingredient) Validate() error {
	switch {
	case i.ID < 1:
		return fmt.Errorf("%w: ingredient id '%d' is not valid", ErrInvalidID, i.ID)
	case len(i.Name) == 0:
		return fmt.Errorf("%w: ingredient name '%s' is not valid", ErrInvalidName, i.Name)
	case len(i.Description) == 0:
		return fmt.Errorf("%w: ingredient description '%s' is not valid", ErrInvalidDescription, i.Description)
	}

	return i.NutritionalInformation.Validate()
}

type NutritionalInformation struct {
	Calories                     int64
	Fats, Proteins, Carbs, Fiber Mass
}

func (i NutritionalInformation) Validate() error {
	return nil
}
