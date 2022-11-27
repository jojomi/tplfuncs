package tplfuncs

import (
	"github.com/spf13/cast"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// CastHelpers returns a text template FuncMap with cast functions
func CastHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"toString": cast.ToString,
	}
}

// CastHelpersHTML returns an HTML template FuncMap with cast functions
func CastHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(CastHelpers())
}

