package tplfuncs

import (
	htmlTemplate "html/template"
	"strings"
	textTemplate "text/template"
)

// SpacingHelpers returns a text template FuncMap with spacing related functions
func SpacingHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"space":     spaceFunc,
		"tab":       tabFunc,
		"newline":   newlineFunc,
		"noop":      noopFunc,
		"blackhole": blackholeFunc,
	}
}

// SpacingHelpersHTML returns an HTML template FuncMap with spacing related functions
func SpacingHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(SpacingHelpers())
}

// Doc: `space` inserts a number of space characters, default is one. Often this function is used in a block that removes whitespace around it.
func spaceFunc(count ...int) string {
	// allows for "space" and "space 4"
	return repeatChar(" ", count...)
}

// Doc: `tab` inserts a number of tab characters, default is one. Often this function is used in a block that removes whitespace around it.
func tabFunc(count ...int) string {
	// allows for "tab" and "tab 4"
	return repeatChar("\t", count...)
}

// Doc: `newline` inserts a number of newline characters, default is one. Often this function is used in a block that removes whitespace around it.
func newlineFunc(count ...int) string {
	// allows for "newline" and "newline 4"
	return repeatChar("\n", count...)
}

func repeatChar(base string, count ...int) string {
	if len(count) == 0 {
		return base
	}
	return strings.Repeat(base, count[0])
}

// Doc: `noop` does nothing. This can be useful to control spacing between elements because {{- -}} is not valid in itself.
func noopFunc() string {
	return ""
}

// Doc: `blackhole` does take any input and discards it.
func blackholeFunc(inputs ...interface{}) string {
	return ""
}
