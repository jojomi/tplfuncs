package tplfuncs

import (
	htmlTemplate "html/template"
	"strings"
	textTemplate "text/template"
)

func SpacingHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"space": func(count ...int) string {
			// allows for "space" and "space 4"
			if len(count) == 0 {
				return " "
			}
			return strings.Repeat(" ", count[0])
		},
		"newline": func(count ...int) string {
			// allows for "newline" and "newline 4"
			if len(count) == 0 {
				return "\n"
			}
			return strings.Repeat("\n", count[0])
		},
	}
}

func SpacingHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(SpacingHelpers())
}
