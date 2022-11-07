package tplfuncs

import (
	"errors"
	"github.com/spf13/afero"
	htmlTemplate "html/template"
	"os"
	textTemplate "text/template"
)

// Fs is the filesystem abstraction to be used
var Fs = afero.NewOsFs()

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
	return afero.Glob(Fs, pattern)
}

func fileExistsFunc(filename string) bool {
	fileInfo, err := Fs.Stat(filename)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return !fileInfo.IsDir()
}

func dirExistsFunc(dirname string) bool {
	fileInfo, err := Fs.Stat(dirname)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return fileInfo.IsDir()
}

func ensureDirFunc(dirname string) error {
	if dirExistsFunc(dirname) {
		return nil
	}

	return Fs.MkdirAll(dirname, 0750)
}

func isMinFileSizeFunc(filename string, minBytes int64) bool {
	if !fileExistsFunc(filename) {
		return false
	}

	fileInfo, err := Fs.Stat(filename)
	if err != nil {
		panic(err)
	}

	return fileInfo.Size() > minBytes
}
