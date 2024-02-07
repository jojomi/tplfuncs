package tplfuncs

import (
	"github.com/cloudfoundry-attic/jibber_jabber"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	htmlTemplate "html/template"
	textTemplate "text/template"
)

// PrintHelpers returns a text template FuncMap with output related functions
func PrintHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"printfLocalized": printfLocalizedFunc,
		"printInt":        printIntFunc,
		"printInt64":      printInt64Func,
	}
}

// PrintHelpersHTML returns an HTML template FuncMap with output related functions
func PrintHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(PrintHelpers())
}

// Doc: `printfLocalized` returns the
func printfLocalizedFunc(key message.Reference, data ...interface{}) (string, error) {
	lang, err := jibber_jabber.DetectLanguage()
	if err != nil {
		return "", err
	}
	p := message.NewPrinter(language.MustParse(lang))
	return p.Sprintf(key, data...), nil
}

// Doc: `printInt` returns the int value as a string.
func printIntFunc(value int) (string, error) {
	return printfLocalizedFunc("%d", value)
}

// Doc: `printInt64` returns the int value as a string.
func printInt64Func(value int64) (string, error) {
	return printfLocalizedFunc("%d", value)
}
