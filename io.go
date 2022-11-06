package tplfuncs

import (
	htmlTemplate "html/template"
	"os"
	textTemplate "text/template"
)

// IOHelpers returns a text template FuncMap with io related functions
func IOHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"readFile":  readFileFunc,
		"writeFile": writeFileFunc,
	}
}

// IOHelpersHTML returns an HTML template FuncMap with io related functions
func IOHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(IOHelpers())
}

func readFileFunc(filename string) (string, error) {
	out, err := os.ReadFile(filename)
	return string(out), err
}

func writeFileFunc(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0640)
}
