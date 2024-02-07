package tplfuncs

import (
	htmlTemplate "html/template"
)

func HTMLSafeHelpers() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap{
		"safeHTML": safeHTMLFunc,
		"safeCSS":  safeCSSFunc,
		"safeJS":   safeJSFunc,
	}
}

// Doc: `safeHTML` declares a string safe to be embedded as HTML without further escaping.
func safeHTMLFunc(input string) htmlTemplate.HTML {
	return htmlTemplate.HTML(input)
}

// Doc: `safeCSS` declares a string safe to be embedded as CSS without further escaping.
func safeCSSFunc(input string) htmlTemplate.CSS {
	return htmlTemplate.CSS(input)
}

// Doc: `safeJS` declares a string safe to be embedded as JavaScript without further escaping.
func safeJSFunc(input string) htmlTemplate.JS {
	return htmlTemplate.JS(input)
}
