{{ with $.msg }} {{- . }}

{{ end -}}

package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"
	"fmt"
)

// DefaultHelpers returns a text template FuncMap with default functions
func DefaultHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
    	{{- range $.defaultTypes -}}
    	    {{- newline -}}
	        {{- $ccName := toCamelCase .name -}}
            {{ $noFirstSet := .no_first_set -}}

	        // functions for {{ $ccName }}
            "firstNonNil{{- $ccName -}}": firstNonNil{{- $ccName -}}Func,
            {{ if $noFirstSet -}}
                "firstSet{{- $ccName -}}": firstNonNil{{- $ccName -}}Func,
                "default{{- $ccName -}}": default{{- $ccName -}}Func,
            {{- else -}}
                "firstSet{{- $ccName -}}": firstSet{{- $ccName -}}Func,
                "default{{- $ccName -}}": default{{- $ccName -}}Func, // alias for firstSet{{- $ccName -}}
            {{ end -}}
    	    {{- newline -}}
        {{ end -}}
	}
}

// DefaultHelpersHTML returns an HTML template FuncMap with default functions
func DefaultHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(DefaultHelpers())
}


{{ range $.defaultTypes -}}
    {{ $ccName := toCamelCase .name -}}
    {{/* $value := firstSet .value .name -*/}}
    {{ $value := .value }}
    {{ if not $value -}}
        {{ $value = .name -}}
    {{ end -}}
    {{ $noFirstSet := .no_first_set -}}

    // Doc: `firstNonNil{{- $ccName -}}` returns the first element in the given list of {{ .name }} values that is not nil.
    func firstNonNil{{- $ccName -}}Func(inputs ...any) ({{ $value }}, error) {
        var empty {{ $value }}

        for _, input := range inputs {
            if input == nil {
                continue
            }

            // is it a {{ $value }}?
            realValue, ok := input.({{- $value -}})

            if !ok {
                return empty, fmt.Errorf("bad: %v (%t)", input, input)
            }

            return realValue, nil
        }
        return empty, fmt.Errorf("all nil!")
    }

    {{ if $noFirstSet -}}
        // Doc: `default{{- $ccName -}}` is an alias for `firstNonNil{{- $ccName -}}`.
        func default{{- $ccName -}}Func(inputs ...any) ({{ $value }}, error) {
            return firstNonNil{{- $ccName -}}Func(inputs...)
        }
    {{ else -}}
        // Doc: `firstSet{{- $ccName -}}` returns the first element in the given list of {{ .name }} values that is not the empty value for {{- $ccName -}}.
        func firstSet{{- $ccName -}}Func(inputs ...any ) (*{{ $value }}, error) {
            var empty {{ $value }}
            for _, input := range inputs {
                var realValue {{ $value }}

                if input == nil {
                    continue
                }

                // is it a {{ $value }} pointer?
                p, ok := input.(*{{- $value -}})
                if ok {
                    if p == nil {
                        continue
                    }
                    realValue = *p
                } else {
                    realValue, ok = input.({{- $value -}})

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

        // Doc: `default{{- $ccName -}}` is an alias for `firstSet{{- $ccName -}}`.
        func default{{- $ccName -}}Func(inputs ...any) (*{{ $value }}, error) {
            return firstSet{{- $ccName -}}Func(inputs...)
        }
    {{ end -}}
{{ end -}}
