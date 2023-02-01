package tplfuncs

import (
	"github.com/iancoleman/strcase"
	htmlTemplate "html/template"
	"regexp"
	"strings"
	textTemplate "text/template"
)

// StringHelpers returns a text template FuncMap with math related functions
func StringHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"stringContains": stringContainsFunc,
		"eqIgnoreCase":   stringEqualFoldFunc,
		"eqFold":         stringEqualFoldFunc,

		// string casing
		"toUpperCase":      stringUpperCaseFunc,
		"toLowerCase":      stringLowerCaseFunc,
		"toCamelCase":      stringCamelCaseFunc,
		"toLowerCamelCase": stringLowerCamelCaseFunc,
		"toSnakeCase":      stringSnakeFunc,
		"toKebabCase":      stringKebabFunc,

		// clean string
		"toCleanString": stringCleanFunc,
		"toFilename":    stringToFilenameFunc,
		"toURL":         stringToURLFunc,
	}
}

// StringHelpersHTML returns an HTML template FuncMap with math related functions
func StringHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(StringHelpers())
}

func stringEqualFoldFunc(a, b string) bool {
	return strings.EqualFold(a, b)
}

func stringContainsFunc(needle, haystack string) bool {
	return strings.Contains(haystack, needle)
}

func stringUpperCaseFunc(input string) string {
	return strings.ToUpper(input)
}

func stringLowerCaseFunc(input string) string {
	return strings.ToLower(input)
}

func stringCamelCaseFunc(input string) string {
	return strcase.ToCamel(input)
}

func stringLowerCamelCaseFunc(input string) string {
	return strcase.ToLowerCamel(input)
}

func stringSnakeFunc(input string) string {
	return strcase.ToSnake(input)
}

func stringKebabFunc(input string) string {
	return strcase.ToKebab(input)
}

func stringCleanFunc(input string) string {
	r := regexp.MustCompile(`[^A-Za-z0-9]+`)
	result := r.ReplaceAllString(input, "_")
	r = regexp.MustCompile(`^_+|_+$`)
	return r.ReplaceAllString(result, "")
}

func stringToFilenameFunc(input string) string {
	return strcase.ToSnake(stringCleanFunc(input))
}

func stringToURLFunc(input string) string {
	return strcase.ToKebab(stringCleanFunc(input))
}
