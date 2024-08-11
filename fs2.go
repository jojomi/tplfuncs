package tplfuncs

import (
	"github.com/jojomi/gofs"
	"github.com/spf13/afero"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// Fs is the filesystem abstraction to be used
var fs = afero.NewOsFs()

// Filesystem2Helpers returns a text template FuncMap with functions related to filesystems
func Filesystem2Helpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"FileAt": func(filepath string) gofs.File {
			return gofs.FileWithFs(filepath, fs)
		},
		"DirAt": func(dirpath string) gofs.Dir {
			return gofs.DirWithFs(dirpath, fs)
		},
	}
}

// Filesystem2HelpersHTML returns an HTML template FuncMap with functions related to filesystems
func Filesystem2HelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(Filesystem2Helpers())
}
