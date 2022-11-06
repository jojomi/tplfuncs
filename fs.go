package tplfuncs

import (
	"errors"
	htmlTemplate "html/template"
	"os"
	"path/filepath"
	textTemplate "text/template"
)

// FilesystemHelpers returns a text template FuncMap with functions related to filesystems
func FilesystemHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"glob":              globFunc,
		"fileExists":        fileExistsFunc,
		"dirExists":         dirExistsFunc,
		"ensureDir":         ensureDirFunc,
		"isMinFileSizeFunc": isMinFileSizeFunc,
	}
}

// FilesystemHelpersHTML returns an HTML template FuncMap with functions related to filesystems
func FilesystemHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(FilesystemHelpers())
}

func globFunc(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

func fileExistsFunc(filename string) bool {
	fileInfo, err := os.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return !fileInfo.IsDir()
}

func dirExistsFunc(dirname string) bool {
	fileInfo, err := os.Stat(dirname)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return fileInfo.IsDir()
}

func ensureDirFunc(dirname string) error {
	if dirExistsFunc(dirname) {
		return nil
	}

	return os.MkdirAll(dirname, 0750)
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
