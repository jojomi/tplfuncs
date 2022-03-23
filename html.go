package tplfuncs

import (
	htmlTemplate "html/template"
)

func HTMLSafeHelpers() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap{
		"safeHTML": func(safeHTMLinput string) htmlTemplate.HTML {
			return htmlTemplate.HTML(safeHTMLinput)
		},
		"safeCSS": func(safeCSSinput string) htmlTemplate.CSS {
			return htmlTemplate.CSS(safeCSSinput)
		},
		"safeJS": func(safeJSinput string) htmlTemplate.JS {
			return htmlTemplate.JS(safeJSinput)
		},
	}
}
