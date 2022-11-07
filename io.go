package tplfuncs

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	htmlTemplate "html/template"
	"os"
	textTemplate "text/template"
)

// IOHelpers returns a text template FuncMap with io related functions
func IOHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"include":   readFileFunc,
		"readFile":  readFileFunc,
		"writeFile": writeFileFunc,
	}
}

// IOHelpersHTML returns an HTML template FuncMap with io related functions
func IOHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(IOHelpers())
}

func readFileFunc(filename string) (string, error) {
	f, err := homedir.Expand(filename)
	if err != nil {
		return "", err
	}

	out, err := afero.ReadFile(Fs, f)
	return string(out), err
}

func writeFileFunc(filename, content string) error {
	return writeFileWithPermsFunc(filename, os.FileMode(0640), content)
}

func writeFileWithPermsFunc(filename string, permissions os.FileMode, content string) error {
	f, err := homedir.Expand(filename)
	if err != nil {
		return err
	}

	return afero.WriteFile(Fs, f, []byte(content), permissions)
}
