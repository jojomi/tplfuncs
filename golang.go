package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/sanity-io/litter"
)

// GolangHelpers returns a text template FuncMap with golang functions
func GolangHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"asGoCode": asGoCodeFunc,
	}
}

// GolangHelpersHTML returns an HTML template FuncMap with golang functions
func GolangHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(GolangHelpers())
}

func asGoCodeFunc(input interface{}) string {
	return litter.Sdump(input)
}
