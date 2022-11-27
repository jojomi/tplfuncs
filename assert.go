package tplfuncs

import (
	"fmt"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// AssertHelpers returns a text template FuncMap with assert functions
func AssertHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"assertString": assertStringFunc,
		"assertInt":    assertIntFunc,
		"assertFloat":  assertFloatFunc,
	}
}

// AssertHelpersHTML returns an HTML template FuncMap with assert functions
func AssertHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(AssertHelpers())
}

func assertTypeFunc[T any](msg string, input interface{}) error {
	_, ok := input.(T)
	if !ok {
		return fmt.Errorf(msg, input, input)
	}
	return nil
}

func assertStringFunc(input interface{}) error {
	msg := fmt.Sprintf("not a string: %+v (type %T)", input, input)
	return assertTypeFunc[string](msg, input)
}

func assertIntFunc(input interface{}) error {
	msg := fmt.Sprintf("not an int: %+v (type %T)", input, input)
	return assertTypeFunc[string](msg, input)
}

func assertFloatFunc(input interface{}) error {
	msg := fmt.Sprintf("not a float: %+v (type %T)", input, input)
	return assertTypeFunc[string](msg, input)
}
