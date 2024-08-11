package tplfuncs

import (
	htmlTemplate "html/template"
	"strconv"
	textTemplate "text/template"
)

// TypeConversionHelpers returns a text template FuncMap with typeConversion functions
func TypeConversionHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"atoi": atoiFunc,
	}
}

// TypeConversionHelpersHTML returns an HTML template FuncMap with typeConversion functions
func TypeConversionHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(TypeConversionHelpers())
}

// Doc: `atoi` returns the integer value in a string.
func atoiFunc(input string) (int, error) {
	return strconv.Atoi(input)
}
