package mystery

// LangType is an enum which represents a programming language
type LangType string

// Enumerators for the language types that are supported. If you are adding a
// new one, please add alphabetically.
const (
	NoLang     LangType = ""
	GoLang              = "go"
	KotlinLang          = "kotlin"
	SwiftLang           = "swift"
	PythonLang          = "python"
)
