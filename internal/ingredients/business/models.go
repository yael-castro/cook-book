package business

import (
	"errors"
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
		return errors.New("missing ingredient id")
	case i.Name == "":
		return errors.New("missing ingredient name")
	case i.Description == "":
		return errors.New("missing ingredient description")
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
