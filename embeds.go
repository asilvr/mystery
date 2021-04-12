package mystery

import (
	_ "embed"
)

//go:embed datasets/go.json
var GoIngredients []byte

//go:embed datasets/kotlin.json
var KotlinIngredients []byte

//go:embed datasets/swift.json
var SwiftIngredients []byte
