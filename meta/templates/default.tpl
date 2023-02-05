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
	        // functions for {{ $ccName }}
            "firstNonNil{{- $ccName -}}": firstNonNil{{- $ccName -}}Function,
            "firstSet{{- $ccName -}}": firstSet{{- $ccName -}}Function,
            "default{{- $ccName -}}": firstSet{{- $ccName -}}Function, // alias for firstSet{{- $ccName -}}
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

    func firstNonNil{{- $ccName -}}Function(inputs ...any) ({{ $value }}, error) {
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

    func firstSet{{- $ccName -}}Function(inputs ...any ) (*{{ $value }}, error) {
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
{{ end -}}
