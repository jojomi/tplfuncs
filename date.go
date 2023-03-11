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
	return htmlTemplate.FuncMap(LineHelpers())
}

func nowFunc() time.Time {
	return time.Now()
}

func parseDateFunc(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}
