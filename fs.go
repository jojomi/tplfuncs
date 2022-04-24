package tplfuncs

import (
	htmlTemplate "html/template"
	"os"
	"path/filepath"
	textTemplate "text/template"
)

// FilesystemHelpers returns a text template FuncMap with functions related to filesystems
func FilesystemHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"glob":    globFunc,
		"include": includeFunc,
	}
}

// FilesystemHTML returns an HTML template FuncMap with functions related to filesystems
func FilesystemHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LineHelpers())
}

func globFunc(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

func includeFunc(filename string) (string, error) {
	b, err := os.ReadFile(filename)
	return string(b), err
}
