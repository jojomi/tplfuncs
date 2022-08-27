package tplfuncs

import (
	htmlTemplate "html/template"
	"io"
	"net/http"
	"os"
	textTemplate "text/template"
)

// NetworkHelpers returns a text template FuncMap with network functions
func NetworkHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"download": downloadFunc,
	}
}

// NetworkHelpersHTML returns an HTML template FuncMap with network functions
func NetworkHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LineHelpers())
}

func downloadFunc(srcURL, filename string) error {
	resp, err := http.Get(srcURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	return err
}
