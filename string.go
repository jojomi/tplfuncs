package tplfuncs

import (
	"fmt"
	htmlTemplate "html/template"
	"regexp"
	"strings"
	textTemplate "text/template"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/iancoleman/strcase"
)

// StringHelpers returns a text template FuncMap with math related functions
func StringHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"trim":            trimFunc,
		"stringContains":  stringContainsFunc,
		"stringHasPrefix": stringHasPrefixFunc,
		"stringHasSuffix": stringHasSuffixFunc,
		"eqIgnoreCase":    stringEqualFoldFunc,
		"eqFold":          stringEqualFoldFunc,

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
		"deHTML":        deHTMLFunc,

		// diff
		"diff": diffFunc,
	}
}

func trimFunc(input string) string {
	return strings.TrimSpace(input)
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

func stringHasPrefixFunc(prefix, testString string) bool {
	return strings.HasPrefix(testString, prefix)
}

func stringHasSuffixFunc(suffix, testString string) bool {
	return strings.HasSuffix(testString, suffix)
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

func diffFunc(nameA, contentA, nameB, contentB string, numContextLines int) string {
	edits := myers.ComputeEdits(span.URIFromPath(nameA), contentA, contentB)
	diff := fmt.Sprint(gotextdiff.ToUnified(nameA, nameB, contentA, edits))
	return diff
}

func deHTMLFunc(input htmlTemplate.HTML) string {
	return string(input)
}
