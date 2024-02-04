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
	f := gofs.FileWithFs(filename, Fs)
	out, err := f.ContentString()
	return out, err
}

func writeFileFunc(filename, content string) error {
	f := gofs.FileWithFs(filename, Fs)
	return f.SetContentString(content)
}

func writeFileWithPermsFunc(filename string, permissions os.FileMode, content string) error {
	f, err := homedir.Expand(filename)
	if err != nil {
		return err
	}

	return afero.WriteFile(Fs, f, []byte(content), permissions)
}
