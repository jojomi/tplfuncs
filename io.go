package tplfuncs

import (
	htmlTemplate "html/template"
	"os"
	textTemplate "text/template"

	"github.com/jojomi/gofs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
)

// IOHelpers returns a text template FuncMap with io related functions
func IOHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"include":            includeFunc,
		"readFile":           readFileFunc,
		"writeFile":          writeFileFunc,
		"writeFileWithPerms": writeFileWithPermsFunc,
	}
}

// IOHelpersHTML returns an HTML template FuncMap with io related functions
func IOHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(IOHelpers())
}

// Doc: `include` is an alias for `readFile`.
func includeFunc(filename string) (string, error) {
	return readFileFunc(filename)
}

// Doc: `readFile` does return the content of a file as a string.
func readFileFunc(filename string) (string, error) {
	f := gofs.FileWithFs(filename, Fs)
	out, err := f.ContentString()
	return out, err
}

// Doc: `writeFile` writes as string to a file.
func writeFileFunc(filename, content string) error {
	f := gofs.FileWithFs(filename, Fs)
	return f.SetContentString(content)
}

// Doc: `writeFileWithPerms` writes as string to a file with given (unix) permissions.
func writeFileWithPermsFunc(filename string, permissions os.FileMode, content string) error {
	f, err := homedir.Expand(filename)
	if err != nil {
		return err
	}

	return afero.WriteFile(Fs, f, []byte(content), permissions)
}
