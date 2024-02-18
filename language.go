package tplfuncs

import (
	"fmt"
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/jojomi/tplfuncs/text"
)

// LanguageHelpers returns a text template FuncMap with functions related to human language
func LanguageHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"plural":             pluralFunc,
		"pluralInt64":        pluralInt64Func,
		"pluralFloat":        pluralFloatFunc,
		"pluralWithNum":      pluralWithNumFunc,
		"pluralInt64WithNum": pluralInt64WithNumFunc,
		"pluralFloatWithNum": pluralFloatWithNumFunc,

		"joinText":         joinTextFunc,
		"joinTextStringer": text.JoinStringer,
	}
}

// LanguageHelpersHTML returns an HTML template FuncMap with functions related to text
func LanguageHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LanguageHelpers())
}

// Doc: `plural` returns the correct string depending on an int value given.
func pluralFunc(singular, plural string, value int) string {
	if value == 1 {
		return singular
	}
	return plural
}

// Doc: `pluralInt64` returns the correct string depending on an int64 value given.
func pluralInt64Func(singular, plural string, value int64) string {
	if value == 1 {
		return singular
	}
	return plural
}

// Doc: `pluralFloat` returns the correct string depending on a float value given.
func pluralFloatFunc(singular, plural string, value float64) string {
	if value == 1 {
		return singular
	}
	return plural
}

// Doc: `pluralWithNum` returns the number and the correct string depending on an int value given.
func pluralWithNumFunc(singular, plural string, value int) string {
	return fmt.Sprintf("%v %s", value, pluralFunc(singular, plural, value))
}

// Doc: `pluralInt64WithNum` returns the number and the correct string depending on an int64 value given.
func pluralInt64WithNumFunc(singular, plural string, value int64) string {
	return fmt.Sprintf("%v %s", value, pluralInt64Func(singular, plural, value))
}

// Doc: `pluralFloatWithNum` returns the number and the correct string depending on a float value given.
func pluralFloatWithNumFunc(singular, plural string, value float64) string {
	return fmt.Sprintf("%v %s", value, pluralFloatFunc(singular, plural, value))
}

// Doc: `joinText` joins elements suitable for a human-readable text.
func joinTextFunc(delim, twoDelim, lastDelim string, input []string) string {
	return text.Join(delim, twoDelim, lastDelim, input)
}
