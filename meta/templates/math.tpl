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
            "divide{{ $ccName }}By": divide{{ $ccName }}ByFunc,
            {{- newline -}}
		{{- end }}
	}
}

// MathHelpersHTML returns an HTML template FuncMap with math related functions
func MathHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(MathHelpers())
}

{{ range $.mathTypes -}}
    {{ $ccName := toCamelCase .name -}}
    {{ $value := .value }}
    {{ if not $value -}}
        {{ $value = .name -}}
    {{ end -}}

    // Doc: `add{{ $ccName }}` adds a number of {{ $value }} values and returns the total sum.
    func add{{ $ccName }}Func(inputs ...{{ $value }}) {{ $value }} {
        var sum {{ $value }}
        for _, input := range inputs {
            sum += input
        }
        return sum
    }

    // Doc: `subtract{{ $ccName }}` subtracts a number of {{ $value }} values from the first one and returns the remaining value.
    func subtract{{ $ccName }}Func(start {{ $value }}, inputs ...{{ $value }}) {{ $value }} {
        sum := start
        for _, input := range inputs {
            sum -= input
        }
        return sum
    }

    // Doc: `subtractFrom{{ $ccName }}` subtracts a number of {{ $value }} values from the last one and returns the remaining value.
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

    // Doc: `multiply{{ $ccName }}` multiplies a number of {{ $value }} values and returns the total value.
    func multiply{{ $ccName }}Func(inputs ...{{ $value }}) {{ $value }} {
        var sum {{ $value }} = 1
        for _, input := range inputs {
            sum *= input
        }
        return sum
    }

    // Doc: `divide{{ $ccName }}By` divides a {{ $value }} value by another one. Note the inverted order to make `24 | divideBy 12` nicely expressive.
    func divide{{ $ccName }}ByFunc(divisor, value {{ $value }}) {{ $value }} {
        return value / divisor
    }
{{ end -}}