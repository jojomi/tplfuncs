package tplfuncs

import (
	"errors"
	"github.com/mitchellh/go-homedir"
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

func fileExistsFunc(filename string) bool {
	fileInfo, err := os.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return !fileInfo.IsDir()
}

func dirExistsFunc(filename string) bool {
	fileInfo, err := os.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return fileInfo.IsDir()
}

func isMinFileSizeFunc(filename string, minBytes int64) bool {
	if !fileExistsFunc(filename) {
		return false
	}

	fileInfo, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	return fileInfo.Size() > minBytes
}

func printFileFunc(filename string) (string, error) {
	f, err := homedir.Expand(filename)
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(f)
	return string(data), err
}
