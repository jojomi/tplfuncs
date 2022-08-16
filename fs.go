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
		"glob":                globFunc,
		"include":             includeFunc,
		"saveToFile":          saveToFileFunc,
		"saveToFileWithPerms": saveToFileWithPermsFunc,
		"fileExists":          fileExistsFunc,
		"dirExists":           dirExistsFunc,
		"isMinFileSizeFunc":   isMinFileSizeFunc,
		"printFile":           printFileFunc,
	}
}

// FilesystemHelpersHTML returns an HTML template FuncMap with functions related to filesystems
func FilesystemHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(FilesystemHelpers())
}

func globFunc(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

func includeFunc(filename string) (string, error) {
	b, err := os.ReadFile(filename)
	return string(b), err
}

func saveToFileFunc(filename, content string) error {
	return saveToFileWithPermsFunc(filename, os.FileMode(0640), content)
}

func saveToFileWithPermsFunc(filename string, permissions os.FileMode, content string) error {
	return os.WriteFile(filename, []byte(content), permissions)
}
