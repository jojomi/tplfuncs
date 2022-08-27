package tplfuncs

import (
	"github.com/cloudfoundry-attic/jibber_jabber"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	htmlTemplate "html/template"
	textTemplate "text/template"
)

// OutputHelpers returns a text template FuncMap with output related functions
func OutputHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"printfLocalized": printfLocalized,
		"printInt": func(value int) (string, error) {
			return printfLocalized("%d", value)
		},
		"printInt64": func(value int64) (string, error) {
			return printfLocalized("%d", value)
		},
	}
}

// OutputHelpersHTML returns an HTML template FuncMap with output related functions
func OutputHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(OutputHelpers())
}

func printfLocalized(key message.Reference, data ...interface{}) (string, error) {
	lang, err := jibber_jabber.DetectLanguage()
	if err != nil {
		return "", err
	}
	p := message.NewPrinter(language.MustParse(lang))
	return p.Sprintf(key, data...), nil
}
