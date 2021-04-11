package mystery

import (
	"fmt"
)

// Ingredient represents a recommended code repository
type Ingredient struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Native      bool   `json:"native"`
	URL         string `json:"url"`
	Import      string `json:"import"`
	Difficulty  string `json:"difficulty"`
}

// Pretty formats the struct for pretty printing
func (i Ingredient) Pretty() string {
	return fmt.Sprintf("%s (%s): %s", i.Name, i.URL, i.Description)
}
