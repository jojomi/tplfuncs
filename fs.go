package tplfuncs

import (
	"errors"
	"fmt"
	htmlTemplate "html/template"
	"os"
	"path"
	"path/filepath"
	"strings"
	textTemplate "text/template"

	"github.com/jojomi/gofs"
	"github.com/spf13/afero"
)

// Fs is the filesystem abstraction to be used
var Fs = afero.NewOsFs()

// FilesystemHelpers returns a text template FuncMap with functions related to filesystems
func FilesystemHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"glob": globFunc,

		"fileAt": fileAtFunc,
		"dirAt":  dirAtFunc,

		"fileExists":        fileExistsFunc,
		"dirExists":         dirExistsFunc,
		"ensureDir":         ensureDirFunc,
		"isMinFileSizeFunc": isMinFileSizeFunc,
		"basename":          basenameFunc,
		"dirname":           dirnameFunc,
		"ext":               extFunc,
		"rawExt":            rawExtFunc,
		"withExt":           withExtFunc,
	}
}

// FilesystemHelpersHTML returns an HTML template FuncMap with functions related to filesystems
func FilesystemHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(FilesystemHelpers())
}

// Doc: `glob` returns a list of files and/or directories matching the patter given.
func globFunc(pattern string) ([]string, error) {
	return afero.Glob(Fs, pattern)
}

// Doc: `fileAt` returns a gofs.File at the given path.
func fileAtFunc(filePath string) gofs.File {
	return gofs.FileWithFs(filePath, Fs)
}

// Doc: `dirAt` returns a gofs.Dir at the given path.
func dirAtFunc(dirPath string) gofs.Dir {
	return gofs.DirWithFs(dirPath, Fs)
}

// Doc: `basename` returns the basename component of a path.
func basenameFunc(filename string) string {
	return filepath.Base(filename)
}

func dirnameFunc(filename string) string {
	return filepath.Dir(filename)
}

func extFunc(filename string) string {
	return filepath.Ext(filename)
}

func rawExtFunc(filename string) string {
	return strings.TrimLeft(filepath.Ext(filename), ".")
}

func withExtFunc(filename, newExt string) string {
	if !strings.HasPrefix(newExt, ".") {
		newExt = "." + newExt
	}
	ext := path.Ext(filename)
	return filename[0:len(filename)-len(ext)] + newExt
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

	return Fs.MkdirAll(dirname, 0o750)
}

func isMinFileSizeFunc(minBytes int64, filename string) (bool, error) {
	if !fileExistsFunc(filename) {
		return false, fmt.Errorf("file not found: %s", filename)
	}

	fileInfo, err := Fs.Stat(filename)
	if err != nil {
		return false, err
	}

	return fileInfo.Size() > minBytes, nil
}
