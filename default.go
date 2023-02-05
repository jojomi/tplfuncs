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
		// functions for Int
		"firstNonNilInt": firstNonNilIntFunction,
		"firstSetInt":    firstSetIntFunction,
		"defaultInt":     firstSetIntFunction, // alias for firstSetInt

		// functions for String
		"firstNonNilString": firstNonNilStringFunction,
		"firstSetString":    firstSetStringFunction,
		"defaultString":     firstSetStringFunction, // alias for firstSetString

		// functions for Float
		"firstNonNilFloat": firstNonNilFloatFunction,
		"firstSetFloat":    firstSetFloatFunction,
		"defaultFloat":     firstSetFloatFunction, // alias for firstSetFloat
	}
}

// DefaultHelpersHTML returns an HTML template FuncMap with default functions
func DefaultHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(DefaultHelpers())
}

func firstNonNilIntFunction(inputs ...any) (int, error) {
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

func firstSetIntFunction(inputs ...any) (*int, error) {
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

func firstNonNilStringFunction(inputs ...any) (string, error) {
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

func firstSetStringFunction(inputs ...any) (*string, error) {
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

func firstNonNilFloatFunction(inputs ...any) (float64, error) {
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

func firstSetFloatFunction(inputs ...any) (*float64, error) {
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