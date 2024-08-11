package tplfuncs

import (
	"fmt"
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
func atoiFunc(input interface{}) (int, error) {
	msg := fmt.Sprintf("not a string: %+v (type %T)", input, input)
	err := assertTypeFunc[string](msg, input)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(input.(string))
}
