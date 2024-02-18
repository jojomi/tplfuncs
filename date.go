package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"
	"time"
)

// DateHelpers returns a text template FuncMap with date functions
func DateHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"now":       nowFunc,
		"parseDate": parseDateFunc,
	}
}

// DateHelpersHTML returns an HTML template FuncMap with date functions
func DateHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(DateHelpers())
}

// Doc: `now` returns the current time (time.Time).
func nowFunc() time.Time {
	return time.Now()
}

// Doc: `parseDate` returns the time.Time associated to the give string when interpreted using the given layout.
func parseDateFunc(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}
