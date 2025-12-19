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
		// string checks
		"stringContains":  stringContainsFunc,
		"stringHasPrefix": stringHasPrefixFunc,
		"stringHasSuffix": stringHasSuffixFunc,
		"eqIgnoreCase":    eqIgnoreCaseFunc,
		"eqFold":          eqFoldFunc,

		// string manipulation
		"trim":                 trimFunc,
		"trimPrefix":           trimPrefixFunc,
		"trimSuffix":           trimSuffixFunc,
		"split":                splitFunc,
		"join":                 joinFunc,
		"cat":                  catFunc,
		"replace":              replaceFunc,
		"regexpReplace":        regexpReplaceFunc,
		"regexpReplaceLiteral": regexpReplaceLiteralFunc,

		// string casing
		"toUpperCase":      toUpperCaseFunc,
		"toLowerCase":      toLowerCaseFunc,
		"toCamelCase":      toCamelCaseFunc,
		"toLowerCamelCase": toLowerCamelCaseFunc,
		"toSnakeCase":      toSnakeCaseFunc,
		"toKebabCase":      toKebabCaseFunc,
		"toTitleCase":      toTitleCaseFunc,

		// clean string
		"toFilename": stringToFilenameFunc,
		"toURL":      stringToURLFunc,
		"deHTML":     deHTMLFunc,

		// diff
		"diff": diffFunc,
	}
}

// StringHelpersHTML returns an HTML template FuncMap with math related functions
func StringHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(StringHelpers())
}

// Doc: `trim` removes all leading and trailing whitespace from the given string. Returns the string unchanged if neither exists.
func trimFunc(input string) string {
	return strings.TrimSpace(input)
}

// Doc: `split` splits the given input string by the given separator.
func splitFunc(input, separator string) []string {
	return strings.Split(input, separator)
}

// Doc: `join` joins the given string array using the given separator.
func joinFunc(separator string, elems []string) string {
	return strings.Join(elems, separator)
}

// Doc: `cat` concatenates the given input strings using the given separator.
func catFunc(elems ...string) string {
	return strings.Join(elems, "")
}

// Doc: `eqFold` compares two strings disregarding their casing.
func eqFoldFunc(a, b string) bool {
	return strings.EqualFold(a, b)
}

// Doc: `eqIgnoreCase` is an alias for `eqFold`.
func eqIgnoreCaseFunc(a, b string) bool {
	return eqFoldFunc(a, b)
}

// Doc: `stringContains` checks if one string is contained in another.
func stringContainsFunc(needle, haystack string) bool {
	return strings.Contains(haystack, needle)
}

// Doc: `stringHasPrefix` determines if a string starts with a given other string.
func stringHasPrefixFunc(prefix, testString string) bool {
	return strings.HasPrefix(testString, prefix)
}

// Doc: `stringHasSuffix` determines if a string ends with a given other string.
func stringHasSuffixFunc(suffix, testString string) bool {
	return strings.HasSuffix(testString, suffix)
}

// Doc: `toUpperCase` returns the given string converted to all uppercase letters.
func toUpperCaseFunc(input string) string {
	return strings.ToUpper(input)
}

// Doc: `toLowerCase` returns the given string converted to all lowercase letters.
func toLowerCaseFunc(input string) string {
	return strings.ToLower(input)
}

// Doc: `toCamelCase` returns the given string converted to camel case (https://en.wikipedia.org/wiki/Camel_case), first letter uppercase.
func toCamelCaseFunc(input string) string {
	return strcase.ToCamel(input)
}

// Doc: `toLowerCamelCase` returns the given string converted to lower camel case (https://en.wikipedia.org/wiki/Camel_case), first letter lowercase.
func toLowerCamelCaseFunc(input string) string {
	return strcase.ToLowerCamel(input)
}

// Doc: `toSnakeCase` returns the given string converted to snake case (https://en.wikipedia.org/wiki/Snake_case).
func toSnakeCaseFunc(input string) string {
	return strcase.ToSnake(input)
}

// Doc: `toKebabCase` returns the given string converted to kebab case (https://en.wikipedia.org/wiki/Kebab_case).
func toKebabCaseFunc(input string) string {
	return strcase.ToKebab(input)
}

// Doc: `toTitleCase` returns the given string converted to title case (https://en.wikipedia.org/wiki/Title_case).
func toTitleCaseFunc(input string) string {
	return strings.ToTitle(input)
}

// Doc: `trimPrefix` returns the given string without the given prefix if there is one, otherwise the string is returned unchanged.
func trimPrefixFunc(prefix, input string) string {
	return strings.TrimPrefix(input, prefix)
}

// Doc: `trimSuffix` returns the given string without the given suffix if there is one, otherwise the string is returned unchanged.
func trimSuffixFunc(suffix, input string) string {
	return strings.TrimSuffix(input, suffix)
}

func stringCleanFunc(replacement rune, input string) string {
	repl := regexp.MustCompile(`[^A-Za-z0-9]+`)
	replacedString := repl.ReplaceAllString(input, string(replacement))

	// remove leading, doubled and trailing replacement runes
	var (
		r                 rune
		result            = ""
		activeReplacement = false
	)
	for _, r = range replacedString {
		if r == replacement {
			// skip doubled
			if activeReplacement {
				continue
			} else {
				// skip leading
				if len(result) == 0 {
					continue
				} else {
					activeReplacement = true
				}
				continue
			}
		}

		// r != replacement
		if activeReplacement {
			result += string(replacement)
			activeReplacement = false
		}

		result += string(r)
	}

	return result
}

// Doc: 'stringToFilename' returns the given string suitable for a filename.
func stringToFilenameFunc(input string) string {
	return strcase.ToSnake(stringCleanFunc('_', input))
}

// Doc: 'stringToURL' returns the given string suitable for a URL.
func stringToURLFunc(input string) string {
	return strcase.ToKebab(stringCleanFunc('-', input))
}

// Doc: `diffFunc` returns the diff between two strings with their associated names.
func diffFunc(nameA, contentA, nameB, contentB string, numContextLines int) string {
	edits := myers.ComputeEdits(span.URIFromPath(nameA), contentA, contentB)
	diff := fmt.Sprint(gotextdiff.ToUnified(nameA, nameB, contentA, edits))
	return diff
}

// Doc: `deHTML` returns the raw string contained in a template.HTML.
func deHTMLFunc(input htmlTemplate.HTML) string {
	return string(input)
}

// Doc: `replace` returns a given string with all occurrences of the given substring replaced by the replacement string.
func replaceFunc(search, replacement, input string) string {
	return strings.ReplaceAll(input, search, replacement)
}

// Doc: `regexpReplace` returns a given string with all occurrences of the given regexp replaced by the replacement string.
func regexpReplaceFunc(regexpValue, replacement, input string) string {
	r := regexp.MustCompile(regexpValue)
	return r.ReplaceAllString(input, replacement)
}

// Doc: `regexpReplaceLiteral` returns a given string with all occurrences of the given regexp replaced by the literal replacement string.
func regexpReplaceLiteralFunc(regexpValue, replacement, input string) string {
	r := regexp.MustCompile(regexpValue)
	return r.ReplaceAllLiteralString(input, replacement)
}

// Doc: `isSingleLine` returns true iff a string does not contain a newline character.
func isSingleLineFunc(input string) bool {
	return !strings.Contains(input, "\n")
}

// Doc: `isMultiline` returns true iff a string does contain a newline character.
func isMultilineFunc(input string) bool {
	return strings.Contains(input, "\n")
}
