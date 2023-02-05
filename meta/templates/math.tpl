{{ with $.msg }} {{- . }}

{{ end -}}

package tplfuncs

import (
	"fmt"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// MathHelpers returns a text template FuncMap with math related functions
func MathHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
    	{{- range $.mathTypes -}}
            {{- newline -}}
            // functions working on {{ .name }} values
            {{ $ccName := toCamelCase .name -}}
            "add{{ $ccName }}": add{{ $ccName }}Func,
            "subtract{{ $ccName }}": subtract{{ $ccName }}Func,
            "subtractFrom{{ $ccName }}": subtractFrom{{ $ccName }}Func,
            "multiply{{ $ccName }}": multiply{{ $ccName }}Func,
            {{- newline -}}
		{{- end }}

        // deprecated
        "floatAdd":   floatAddFunc,
        "floatSub":   floatSubFunc,
        "floatMul":   floatMulFunc,
        "floatDiv":   floatDivFunc,
        "floatDivBy": floatDivByFunc,
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

func floatSubFunc(values ...float64) float64 {
	if len(values) == 0 {
		return 0.0
	}

	sum := values[0]
	for _, v := range values[1:] {
		sum -= v
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


{{ range $.mathTypes -}}
    {{ $ccName := toCamelCase .name -}}
    {{ $value := .value }}
    {{ if not $value -}}
        {{ $value = .name -}}
    {{ end -}}

    // add{{ $ccName }}Func adds a number of {{ $value }} values and returns the total sum.
    func add{{ $ccName }}Func(inputs ...{{ $value }}) {{ $value }} {
        var sum {{ $value }}
        for _, input := range inputs {
            sum += input
        }
        return sum
    }

    // subtract{{ $ccName }}Func subtracts a number of {{ $value }} values from the first one and returns the remaining value.
    func subtract{{ $ccName }}Func(start {{ $value }}, inputs ...{{ $value }}) {{ $value }} {
        sum := start
        for _, input := range inputs {
            sum -= input
        }
        return sum
    }

    // subtractFrom{{ $ccName }}Func subtracts a number of {{ $value }} values from the last one and returns the remaining value.
    func subtractFrom{{ $ccName }}Func(inputs ...{{ $value }}) {{ $value }} {
        if len(inputs) == 0 {
            return 0
        }
        sum := inputs[len(inputs)-1]
        for i := 0; i < len(inputs)-1; i++ {
            sum -= inputs[i]
        }
        return sum
    }

    // multiply{{ $ccName }}Func multiplies a number of {{ $value }} values and returns the total value.
    func multiply{{ $ccName }}Func(inputs ...{{ $value }}) {{ $value }} {
        var sum {{ $value }}
        for _, input := range inputs {
            sum *= input
        }
        return sum
    }
{{ end -}}