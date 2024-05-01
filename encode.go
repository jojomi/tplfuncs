package tplfuncs

import (
	htmlTemplate "html/template"
	"net/url"
	textTemplate "text/template"
)

// EncodeHelpers returns a text template FuncMap with encode functions
func EncodeHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"urlencode": urlencodeFunc,
		"urldecode": urldecodeFunc,
	}
}

// EncodeHelpersHTML returns an HTML template FuncMap with encode functions
func EncodeHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LinesHelpers())
}

// Doc: `urlencode` returns urlencoded value of the input.
func urlencodeFunc(input string) string {
	return url.QueryEscape(input)
}

// Doc: `urldecode` returns the urldecoded value of the input.
func urldecodeFunc(input string) (string, error) {
	decoded, err := url.QueryUnescape(input)
	if err != nil {
		return "", err
	}
	return decoded, nil
}
