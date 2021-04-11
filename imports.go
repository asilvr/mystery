package mystery

import "fmt"

// ImportStatement generates an import statement based on the language and the
// ingredients.
func ImportStatement(language LangType, ingredients []Ingredient) string {
	statement := ""
	switch language {
	case GoLang:
		statement = genImportGo(ingredients)
	case SwiftLang:
		statement = genImportSwift(ingredients)
	}

	return statement
}

// genImportGo is a helper to generate Go import statements
func genImportGo(ingredients []Ingredient) string {
	statement := ""
	for _, ingredient := range ingredients {
		statement = fmt.Sprintf("%s\t\"%s\"\n", statement, ingredient.Import)
	}
	statement = fmt.Sprintf("import (\n%s)\n", statement)
	return statement
}

// genImportSwift is a helper to generate Swift import statements
func genImportSwift(ingredients []Ingredient) string {
	statement := ""
	for _, ingredient := range ingredients {
		statement = fmt.Sprintf("%simport %s\n", statement, ingredient.Import)
	}
	return statement
}
