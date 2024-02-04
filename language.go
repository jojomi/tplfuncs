package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/jojomi/tplfuncs/text"
)

// LanguageHelpers returns a text template FuncMap with functions related to human language
func LanguageHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"plural":      pluralFunc,
		"pluralInt64": pluralInt64Func,
		"pluralFloat": pluralFloatFunc,

		"joinText":         text.Join,
		"joinTextStringer": text.JoinStringer,
	}
}

// LanguageHelpersHTML returns an HTML template FuncMap with functions related to text
func LanguageHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LanguageHelpers())
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
