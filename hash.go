package tplfuncs

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// HashHelpers returns a text template FuncMap with hash related functions
func HashHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"sha1":   sha1Func,
		"sha256": sha256Func,
	}
}

// HashHelpersHTML returns an HTML template FuncMap with hash related functions
func HashHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(HashHelpers())
}

func sha1Func(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	sha := h.Sum(nil)

	return hex.EncodeToString(sha)
}

func sha256Func(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	sha := h.Sum(nil)

	return hex.EncodeToString(sha)
}
