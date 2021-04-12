package mystery

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// Dataset is a collection of ingredients
type Dataset struct {
	Ingredients []Ingredient `json:"ingredients"`
}

// LoadDataset loads an ingredients dataset from an embedded file into a struct
func LoadDataset(language LangType) (ds Dataset, _ error) {
	raw := []byte{}

	switch language {
	case GoLang:
		raw = GoIngredients
	case KotlinLang:
		raw = KotlinIngredients
	case SwiftLang:
		raw = SwiftIngredients
	default:
		return ds, fmt.Errorf("language does not have ingredients dataset")
	}

	if err := json.Unmarshal(raw, &ds); err != nil {
		return ds, err
	}

	return ds, nil
}

// Generate generates random mystery ingredients from the dataset
func (d Dataset) Generate(difficulty string, imports string) []Ingredient {
	rand.Seed(time.Now().Unix())

	rand.Shuffle(len(d.Ingredients), func(i, j int) {
		d.Ingredients[i], d.Ingredients[j] = d.Ingredients[j], d.Ingredients[i]
	})

	// TODO: Return data that is considerate of difficult and imports arguments
	return d.Ingredients[0:3]
}
