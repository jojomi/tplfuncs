package tplfuncs

import (
	"github.com/jojomi/tplfuncs/text"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// TextHelpers returns a text template FuncMap with functions related to text
func TextHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"plural":      pluralFunc,
		"pluralInt64": pluralInt64Func,
		"pluralFloat": pluralFloatFunc,

		"joinText":         text.Join,
		"joinTextStringer": text.JoinStringer,
	}
}

// TextHelpersHTML returns an HTML template FuncMap with functions related to text
func TextHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(TextHelpers())
}

func pluralFunc(value int, singular, plural string) string {
	if value == 1 {
		return singular
	}
	return plural
}

func pluralInt64Func(value int64, singular, plural string) string {
	if value == 1 {
		return singular
	}
	return plural
}

func pluralFloatFunc(value float64, singular, plural string) string {
	if value == 1 {
		return singular
	}
	return plural
}
