package tplfuncs

import (
	htmlTemplate "html/template"
	"strconv"
	textTemplate "text/template"
)

// TypeConversionHelpers returns a text template FuncMap with typeConversion functions
func TypeConversionHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"atoi":    atoiFunc,
		"float64": float64Func,
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

// Doc: `float64` returns the float value in a string.
func float64Func(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}
