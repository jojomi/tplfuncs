package tplfuncs

import (
	"fmt"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// MathHelpers returns a text template FuncMap with math related functions
func MathHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"floatDiv":   floatDivFunc,
		"floatDivBy": floatDivByFunc,
		"floatAdd":   floatAddFunc,
		"floatMul":   floatMulFunc,
	}
}

// MathHelpersHTML returns an HTML template FuncMap with math related functions
func MathHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(MathHelpers())
}

func floatDivFunc(values ...float64) (float64, error) {
	if len(values) < 2 {
		return 0, fmt.Errorf("not enough values given for floating point division: %v", values)
	}
	result := values[0]
	for _, v := range values[1:] {
		if v == 0 {
			return 0, fmt.Errorf("floating point division by null with values %v", values)
		}
		result = result / v
	}
	return result, nil
}

func floatDivByFunc(values ...float64) (float64, error) {
	count := len(values)
	reversedValues := make([]float64, count)

	for i, v := range values {
		reversedValues[count-(i+1)] = v
	}

	return floatDivFunc(reversedValues...)
}

func floatAddFunc(values ...float64) float64 {
	if len(values) == 0 {
		return 0.0
	}

	sum := values[0]
	for _, v := range values[1:] {
		sum += v
	}

	return sum
}

func floatMulFunc(values ...float64) float64 {
	if len(values) == 0 {
		return 0.0
	}

	sum := values[0]
	for _, v := range values[1:] {
		sum *= v
	}

	return sum
}
