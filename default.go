// AUTOGENERATED FILE. DO NOT EDIT.

package tplfuncs

import (
	"fmt"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// DefaultHelpers returns a text template FuncMap with default functions
func DefaultHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		// functions for Bool
		"firstNonNilBool": firstNonNilBoolFunc,
		"firstSetBool":    firstNonNilBoolFunc,
		"defaultBool":     defaultBoolFunc,

		// functions for Int
		"firstNonNilInt": firstNonNilIntFunc,
		"firstSetInt":    firstSetIntFunc,
		"defaultInt":     defaultIntFunc, // alias for firstSetInt

		// functions for String
		"firstNonNilString": firstNonNilStringFunc,
		"firstSetString":    firstSetStringFunc,
		"defaultString":     defaultStringFunc, // alias for firstSetString

		// functions for Float
		"firstNonNilFloat": firstNonNilFloatFunc,
		"firstSetFloat":    firstSetFloatFunc,
		"defaultFloat":     defaultFloatFunc, // alias for firstSetFloat
	}
}

// DefaultHelpersHTML returns an HTML template FuncMap with default functions
func DefaultHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(DefaultHelpers())
}

// Doc: `firstNonNilBool` returns the first element in the given list of bool values that is not nil.
func firstNonNilBoolFunc(inputs ...any) (bool, error) {
	var empty bool

	for _, input := range inputs {
		if input == nil {
			continue
		}

		// is it a bool?
		realValue, ok := input.(bool)

		if !ok {
			return empty, fmt.Errorf("bad: %v (%t)", input, input)
		}

		return realValue, nil
	}
	return empty, fmt.Errorf("all nil!")
}

// Doc: `defaultBool` is an alias for `firstNonNilBool`.
func defaultBoolFunc(inputs ...any) (bool, error) {
	return firstNonNilBoolFunc(inputs...)
}

// Doc: `firstNonNilInt` returns the first element in the given list of int values that is not nil.
func firstNonNilIntFunc(inputs ...any) (int, error) {
	var empty int

	for _, input := range inputs {
		if input == nil {
			continue
		}

		// is it a int?
		realValue, ok := input.(int)

		if !ok {
			return empty, fmt.Errorf("bad: %v (%t)", input, input)
		}

		return realValue, nil
	}
	return empty, fmt.Errorf("all nil!")
}

// Doc: `firstSetInt` returns the first element in the given list of int values that is not the empty value forInt.
func firstSetIntFunc(inputs ...any) (*int, error) {
	var empty int
	for _, input := range inputs {
		var realValue int

		if input == nil {
			continue
		}

		// is it a int pointer?
		p, ok := input.(*int)
		if ok {
			if p == nil {
				continue
			}
			realValue = *p
		} else {
			realValue, ok = input.(int)

			if !ok {
				return nil, fmt.Errorf("bad: %v (%t)", input, input)
			}
		}

		if realValue != empty {
			return &realValue, nil
		}
	}
	return nil, nil
}

// Doc: `defaultInt` is an alias for `firstSetInt`.
func defaultIntFunc(inputs ...any) (*int, error) {
	return firstSetIntFunc(inputs...)
}

// Doc: `firstNonNilString` returns the first element in the given list of string values that is not nil.
func firstNonNilStringFunc(inputs ...any) (string, error) {
	var empty string

	for _, input := range inputs {
		if input == nil {
			continue
		}

		// is it a string?
		realValue, ok := input.(string)

		if !ok {
			return empty, fmt.Errorf("bad: %v (%t)", input, input)
		}

		return realValue, nil
	}
	return empty, fmt.Errorf("all nil!")
}

// Doc: `firstSetString` returns the first element in the given list of string values that is not the empty value forString.
func firstSetStringFunc(inputs ...any) (*string, error) {
	var empty string
	for _, input := range inputs {
		var realValue string

		if input == nil {
			continue
		}

		// is it a string pointer?
		p, ok := input.(*string)
		if ok {
			if p == nil {
				continue
			}
			realValue = *p
		} else {
			realValue, ok = input.(string)

			if !ok {
				return nil, fmt.Errorf("bad: %v (%t)", input, input)
			}
		}

		if realValue != empty {
			return &realValue, nil
		}
	}
	return nil, nil
}

// Doc: `defaultString` is an alias for `firstSetString`.
func defaultStringFunc(inputs ...any) (*string, error) {
	return firstSetStringFunc(inputs...)
}

// Doc: `firstNonNilFloat` returns the first element in the given list of float values that is not nil.
func firstNonNilFloatFunc(inputs ...any) (float64, error) {
	var empty float64

	for _, input := range inputs {
		if input == nil {
			continue
		}

		// is it a float64?
		realValue, ok := input.(float64)

		if !ok {
			return empty, fmt.Errorf("bad: %v (%t)", input, input)
		}

		return realValue, nil
	}
	return empty, fmt.Errorf("all nil!")
}

// Doc: `firstSetFloat` returns the first element in the given list of float values that is not the empty value forFloat.
func firstSetFloatFunc(inputs ...any) (*float64, error) {
	var empty float64
	for _, input := range inputs {
		var realValue float64

		if input == nil {
			continue
		}

		// is it a float64 pointer?
		p, ok := input.(*float64)
		if ok {
			if p == nil {
				continue
			}
			realValue = *p
		} else {
			realValue, ok = input.(float64)

			if !ok {
				return nil, fmt.Errorf("bad: %v (%t)", input, input)
			}
		}

		if realValue != empty {
			return &realValue, nil
		}
	}
	return nil, nil
}

// Doc: `defaultFloat` is an alias for `firstSetFloat`.
func defaultFloatFunc(inputs ...any) (*float64, error) {
	return firstSetFloatFunc(inputs...)
}
