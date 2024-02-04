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

// IngredientFilter contains parameters to make ingredient searches
type IngredientFilter struct {
	Page  uint64
	Size  uint64
	Total uint64
	// Keyword is used to compare if match partially with an ingredient name
	Keyword string
	// Random indicates if the ingredients will be searched randomly.
	Random bool
}

func (f IngredientFilter) Validate() error {
	const minSize, maxSize = 1, 25

	if f.Size < minSize && f.Size > maxSize {
		return fmt.Errorf("%w: '%d' is not a valid page size", ErrInvalidFilter, f.Size)
	}

	if !f.Random && len(f.Keyword) == 0 {
		return fmt.Errorf("%w: the random or keyword parameter is required", ErrInvalidFilter)
	}

	return nil
}

type Ingredient struct {
	NutritionalInformation
	ID                int64
	Category          string
	Name, Description string
}

func (i Ingredient) Validate() error {
	switch {
	case i.ID < 1:
		return fmt.Errorf("%w: ingredient id '%d' is not valid", ErrInvalidID, i.ID)
	case len(i.Category) == 0:
		return fmt.Errorf("%w: ingredient category '%s' is not valid", ErrInvalidCategory, i.Category)
	case len(i.Name) == 0:
		return fmt.Errorf("%w: ingredient name '%s' is not valid", ErrInvalidName, i.Name)
	case len(i.Description) == 0:
		return fmt.Errorf("%w: ingredient description '%s' is not valid", ErrInvalidDescription, i.Description)
	}

	return nil
}

type NutritionalInformation struct {
	Calories                     int64
	Fats, Proteins, Carbs, Fiber Mass
}
