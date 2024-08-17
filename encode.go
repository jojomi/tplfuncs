package tplfuncs

import (
	htmlTemplate "html/template"
	"net/url"
	"strings"
	textTemplate "text/template"
)

// EncodeHelpers returns a text template FuncMap with encode functions
func EncodeHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"urlencode":   urlencodeFunc,
		"urldecode":   urldecodeFunc,
		"escapeChars": escapeCharsFunc,
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

// Doc: `escapeChars` returns the input with any of the given chars escaped with the escaping char.
func escapeCharsFunc(input, escapeChars, escapingChar string) string {
	result := input
	for _, char := range input {
		if strings.ContainsRune(escapeChars, char) {
			result += escapingChar + string(char)
		} else {
			result += string(char)
		}
	}
	return result
}
