package mystery

import (
	"fmt"
	"strings"
)

// PrintGenerate prints out the CLI text for the generate command
func PrintGenerate(language string, ingredients []Ingredient) {
	intro := "\nVoila! Your mystery ingredients have been generated! You'll be making a project in the %s language with the following ingredients:\n\n"
	fmt.Printf(intro, language)
	for _, ingredient := range ingredients {
		fmt.Printf("- %s\n", ingredient.Pretty())
	}

	snippet := "\nUse this snippet if you want to import it:\n\n"
	fmt.Printf(snippet)

	fmt.Printf(ImportStatement(LangType(strings.ToLower(language)), ingredients))

	outro := "Good luck on your mystery project!"
	fmt.Printf("\n%s\n\n", outro)
}
